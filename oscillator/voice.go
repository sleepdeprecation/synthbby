package oscillator

type Voice struct {
	// SampleRate float64

	Source        *Oscillator
	PreAmpFilters []PreAmpFilter
	Envelope
	// Amp            Amplifier
	PostAmpFilters []PostAmpFilter
}

func NewSimpleVoice(sampleRate float64) *Voice {
	voice := &Voice{
		Source: &Oscillator{
			waveMultiplier: Tau / sampleRate,
			WaveFn:         downSawWave,
			// WaveFn: sineWave,
		},
		Amp: BasicAmp,
	}

	return voice
}

func (v *Voice) SetFrequency(frequency float64) {
	v.Source.SetFrequency(frequency)
}

func (v *Voice) GetSamples(n int) []int16 {
	frames := make([]float64, n)
	for i := 0; i < n; i++ {
		frames[i] = v.Source.Tick()
	}

	for _, filter := range v.PreAmpFilters {
		filter.Filter(frames)
	}

	samples := make([]int16, n)
	for i := 0; i < n; i++ {
		samples[i] = v.Amp.Amplify(frames[i])
	}

	for _, filter := range v.PostAmpFilters {
		filter.Filter(samples)
	}

	return samples
}

// func (v *Voice) GetSample() int16 {
// 	frame := v.Source.Tick()
// 	for _, filter := range v.PreAmpFilters {
// 		frame = filter.Filter(frame)
// 	}

// 	sample := v.Amp.Amplify(frame)

// 	for _, filter := range v.PostAmpFilters {
// 		sample = filter.Filter(sample)
// 	}

// 	return sample
// }

func (v *Voice) SampleBytes() []byte {
	sample := v.GetSample()
	return []byte{
		byte(sample),
		byte(sample >> 8),
		byte(sample),
		byte(sample >> 8),
	}
}
