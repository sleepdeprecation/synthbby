package synth

import (
	"github.com/sleepdeprecation/synthbby/oscillator"
)

type Synth struct {
	SampleRate int
	Envelope   *Envelope
	Oscillator *oscillator.Oscillator
}

func New(sampleRate int, wave oscillator.WaveFunction) *Synth {
	return &Synth{
		SampleRate: sampleRate,
		Oscillator: oscillator.New(sampleRate, wave),
		Envelope: &Envelope{
			Attack:  0.0,
			Decay:   0.0,
			Sustain: 1.0,
			Release: 0.0,
		},
	}
}

func (s *Synth) BuildStep(numFrames int, pitch float64, filters []Filter, gateOff, gateOpen, gateClose bool) []float64 {
	frames := s.Envelope.BuildStep(numFrames, gateOff, gateOpen, gateClose)

	s.Oscillator.SetFrequency(pitch)
	if gateOpen {
		s.Oscillator.Reset()
	}

	ticks := s.Oscillator.Ticks(numFrames)
	for _, filter := range filters {
		filter(ticks)
	}

	for i := 0; i < numFrames; i++ {
		frames[i] = frames[i] * ticks[i]
	}

	return frames
}
