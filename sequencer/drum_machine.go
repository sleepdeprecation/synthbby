package sequencer

import (
	"math"

	"github.com/sleepdeprecation/synthbby/oscillator"
	"github.com/sleepdeprecation/synthbby/synth"
)

type DrumMachine struct {
	sampleRate int
	clockRate  int

	Snare   *Instrument
	Bass    *Instrument
	Lead    *Instrument
	Counter *Instrument
}

func NewDrumMachine(sampleRate, clockRate int) *DrumMachine {
	snare := &Instrument{
		Sequencer: NewStepSequencer(sampleRate, clockRate),
		Synth:     synth.New(sampleRate, oscillator.NoiseWave),
		Filters:   []synth.Filter{synth.Mute},
	}

	bass := &Instrument{
		Sequencer: NewStepSequencer(sampleRate, clockRate),
		Synth:     synth.New(sampleRate, oscillator.DownSawWave),
		Filters: []synth.Filter{
			synth.Mute,
		},
	}

	lead := &Instrument{
		Sequencer: NewStepSequencer(sampleRate, clockRate),
		Synth:     synth.New(sampleRate, oscillator.SineWave),
		Filters: []synth.Filter{
			synth.WaveFoldBacker(0.9, 0.6),
			// synth.Overdrive(13),
			// synth.BitCrusher(4),
		},
	}

	counter := &Instrument{
		Sequencer: NewStepSequencer(sampleRate, clockRate),
		Synth:     synth.New(sampleRate, oscillator.SineWave),
	}

	machine := &DrumMachine{
		sampleRate: sampleRate,
		clockRate:  clockRate,
		Snare:      snare,
		Bass:       bass,
		Lead:       lead,
		Counter:    counter,
	}

	return machine
}

func (d *DrumMachine) Render() []int16 {
	snare := d.Snare.Render()
	bass := d.Bass.Render()
	lead := d.Lead.Render()
	counter := d.Counter.Render()

	samples := make([]int16, len(snare))
	for i := 0; i < len(snare); i++ {
		samples[i] = int16(((snare[i] + bass[i] + lead[i] + counter[i]) / 4) * math.MaxInt16)
		// samples[i] = int16(((snare[i] + bass[i] + lead[i] + counter[i]) / 4) * math.MaxInt16)
		// samples[i] = int16(int16(float64((snare[i]+bass[i]+lead[i]+counter[i])/4)*16) * 258)
	}

	return samples
}
