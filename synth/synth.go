package synth

import (
	"math"

	"github.com/sleepdeprecation/synthbby/oscillator"
)

type Synth struct {
	SampleRate int
	Envelope   *Envelope
	Oscillator *oscillator.Oscillator
}

func (s *Synth) BuildStep(numFrames int, pitch float64, gateOpen, gateClose bool) []int16 {
	ampData := s.Envelope.BuildStep(numFrames, gateOpen, gateClose)
	stepData := make([]int16, numFrames)

	s.Oscillator.SetFrequency(pitch)
	if gateOpen {
		s.Oscillator.Reset()
	}

	for i := 0; i < numFrames; i++ {
		frame := s.Oscillator.Tick()
		stepData[i] = int16(frame * math.MaxInt16 * ampData[i])
		// stepData[i] = int16(frame * math.MaxInt16) // * ampData[i])
	}

	return stepData
}
