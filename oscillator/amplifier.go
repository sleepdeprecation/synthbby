package oscillator

import (
	"fmt"
	"math"
)

type Amplifier interface {
	// Amplify takes one frame from an oscillator and amplifies it into a 16 bit linear PCM sample
	Amplify(frame float64) int16
}

type ConstantAmp struct {
	Level float64
}

func (a *ConstantAmp) Amplify(frame float64) int16 {
	if a.Level < 0.0 || a.Level > 1.0 {
		panic(fmt.Errorf("amplifier level must be between 0 and 1"))
	}

	return int16(frame * math.MaxInt16 * a.Level)
}

var (
	BasicAmp = &ConstantAmp{Level: 1.0}
)

type OscillatingAmp struct {
	Source *Oscillator
	Amount float64
}

func (a *OscillatingAmp) Amplify(frame float64) int16 {
	ampLevel := a.Source.Tick()
	ampLevel += 1.0           // switch range from -1:1 to 0:2
	ampLevel = ampLevel / 2.0 // switch range to 0:1
	ampLevel = 1 - (ampLevel * a.Amount)

	sample := int16(frame * math.MaxInt16 * ampLevel)
	return sample
}

type PreRenderedAmp struct {
	Samples  []float64
	position int
}

func (a *PreRenderedAmp) Amplify(frame float64) int16 {
	ampLevel := a.Samples[a.position]
	a.position += 1
	a.position %= len(a.Samples)

	return int16(frame * math.MaxInt16 * ampLevel)
}
