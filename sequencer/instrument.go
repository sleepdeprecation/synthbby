package sequencer

import (
	"github.com/sleepdeprecation/synthbby/synth"
)

type Instrument struct {
	Sequencer *StepSequencer
	Synth     *synth.Synth
	Filters   []synth.Filter

	framesPerStep int
}

func (i *Instrument) SetEnvelope(a, d, s, r float64) {
	i.Synth.Envelope.Attack = a
	i.Synth.Envelope.Decay = d
	i.Synth.Envelope.Sustain = s
	i.Synth.Envelope.Release = r
}

func (i *Instrument) Steps() [16]*Step {
	return i.Sequencer.Steps
}

func (i *Instrument) SetStep(n int, pitch float64, gate Gate) {
	i.Sequencer.Steps[n].Pitch = pitch
	i.Sequencer.Steps[n].Gate = gate
}

func (i *Instrument) Render() []float64 {
	if i.framesPerStep == 0 {
		i.framesPerStep = i.Sequencer.SampleRate / (i.Sequencer.ClockRate / 60)
	}

	samples := make([]float64, i.framesPerStep*16)

	for idx, step := range i.Sequencer.Steps {
		stepSamples := i.Synth.BuildStep(
			i.framesPerStep,
			step.Pitch,
			i.Filters,
			step.Gate.IsOff(),
			step.Gate.IsStart(),
			step.Gate.IsEnd(),
		)

		start := idx * i.framesPerStep
		end := (idx + 1) * i.framesPerStep
		copy(samples[start:end], stepSamples)
	}

	return samples
}
