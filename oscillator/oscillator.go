package oscillator

import (
	"math"
)

const (
	Tau = 2 * math.Pi
)

type WaveFunction func(float64) float64

type Oscillator struct {
	frequency      float64
	phase          float64
	incrementer    float64
	waveMultiplier float64
	waveFn         WaveFunction
}

// New set to a given sample rate
func New(sampleRate int, wave WaveFunction) *Oscillator {
	return &Oscillator{
		waveMultiplier: Tau / float64(sampleRate),
		waveFn:         wave,
	}
}

func (o *Oscillator) SetWave(wave WaveFunction) {
	o.waveFn = wave
	o.Reset()
}

func (o *Oscillator) SetFrequency(freq float64) {
	o.frequency = freq
	o.incrementer = o.waveMultiplier * o.frequency
}

func (o *Oscillator) Reset() {
	o.phase = Tau
}

func (o *Oscillator) Tick() float64 {
	tick := o.waveFn(o.phase)
	o.phase += o.incrementer

	if o.phase >= Tau {
		o.phase -= Tau
	}

	if o.phase < 0 {
		o.phase = Tau
	}

	return tick
}

func (o *Oscillator) Ticks(n int) []float64 {
	ticks := make([]float64, n)
	for idx := range ticks {
		ticks[idx] = o.Tick()
	}

	return ticks
}
