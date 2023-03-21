package sequencer

type Note struct {
	// Frequency is the frequency for the oscillator to oscillate at
	Frequency float64

	// Envelope is an ADSR envelope controlling amp information for shaping a sound
	Envelope *Envelope

	// Duration is in number of beats
	Duration float64
}

func Rest(duration float64) *Note {
	return &Note{
		Duration: duration,
		Envelope: EnvelopeZero,
	}
}
