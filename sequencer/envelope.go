package sequencer

type Envelope struct {
	// Attack, Decay, and Release are lengths in beats
	Attack  float64
	Decay   float64
	Release float64

	// Sustain is an amplitude level, and must be between 0 and 1
	Sustain float64
}

var EnvelopeOff = &Envelope{
	Attack:  0.0,
	Decay:   0.0,
	Sustain: 1.0,
	Release: 0.0,
}

var EnvelopeZero = &Envelope{
	Attack:  0.0,
	Decay:   0.0,
	Sustain: 0.0,
	Release: 0.0,
}

// Render returns an "amp slice", or amplifier data for shaping, given a sample rate, a duration (in beats), and a beats per minute rate.
// The amp slice will be a slice of float64 between 0 and 1, and should act as a multiplier to be passed into the amplification stage.
// The length of the returned amp slice will be:
//
//	beats per second = bpm / 60
//	samples per beat = sample rate / beats per second
//	length = duration * samples per beat
//
// or, put together
//
//	length = sampleRate / (bpm / 60) * duration
func (e *Envelope) Render(sampleRate int64, duration float64, bpm int64) []float64 {
	beatsPerSecond := float64(bpm) / 60.0
	samplesPerBeat := float64(sampleRate) / beatsPerSecond

	totalLength := int(samplesPerBeat * duration)

	attackLength := int(samplesPerBeat * e.Attack)
	decayLength := int(samplesPerBeat * e.Decay)

	releaseLength := int(samplesPerBeat * e.Release)

	adLength := attackLength + decayLength
	if (adLength + releaseLength) > totalLength {
		releaseLength = totalLength - adLength
	}

	sustainLength := totalLength - (adLength + releaseLength)

	samples := make([]float64, totalLength)

	// add attack data
	for i := 0; i < attackLength; i++ {
		pos := i

		samples[pos] = float64(i) / float64(attackLength)
	}

	// add decay data
	for i := 0; i < decayLength; i++ {
		pos := attackLength + i

		samples[pos] = 1 - ((float64(i) / float64(decayLength)) * (1 - e.Sustain))
	}

	// add sustain data
	for i := 0; i < sustainLength; i++ {
		pos := attackLength + decayLength + i

		samples[pos] = e.Sustain
	}

	// add release data
	for i := 0; i < releaseLength; i++ {
		pos := attackLength + decayLength + sustainLength + i

		releasePercent := 1 - (float64(i) / float64(releaseLength))
		samples[pos] = releasePercent * e.Sustain
	}

	return samples
}
