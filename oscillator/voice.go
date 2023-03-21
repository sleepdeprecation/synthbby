package oscillator

type Voice struct {
	// SampleRate float64

	Source         *Oscillator
	PreAmpFilters  []PreAmpFilter
	Amp            Amplifier
	PostAmpFilters []PostAmpFilter
}

func NewSimpleVoice(sampleRate float64) *Voice {
	voice := &Voice{
		Source: &Oscillator{
			waveMultiplier: Tau / sampleRate,
			WaveFn:         noiseWave,
		},
		Amp: BasicAmp,
	}

	return voice
}

func (v *Voice) SetFrequency(frequency float64) {
	v.Source.SetFrequency(frequency)
}

func (v *Voice) GetSample() int16 {
	frame := v.Source.Tick()
	for _, filter := range v.PreAmpFilters {
		frame = filter.Filter(frame)
	}

	sample := v.Amp.Amplify(frame)

	for _, filter := range v.PostAmpFilters {
		sample = filter.Filter(sample)
	}

	return sample
}

func (v *Voice) SampleBytes() []byte {
	sample := v.GetSample()
	return []byte{
		byte(sample),
		byte(sample >> 8),
		byte(sample),
		byte(sample >> 8),
	}
}
