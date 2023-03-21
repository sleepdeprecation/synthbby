package oscillator

import (
	"fmt"
	"math"
	"math/rand"
)

type Shape int

const (
	Sine Shape = iota
	Square
	DownSaw
	UpSaw
	Triangle
	Noise
)

const (
	Tau = 2 * math.Pi
)

type WaveFunction func(float64) float64

var (
	shapeCalcFunc = map[Shape]WaveFunction{
		Sine:     sineWave,
		Square:   squareWave,
		Triangle: triangleWave,
		DownSaw:  downSawWave,
		UpSaw:    upSawWave,
		Noise:    noiseWave,
	}
)

type Oscillator struct {
	frequency      float64
	phase          float64
	incrementer    float64
	waveMultiplier float64
	WaveFn         WaveFunction
}

// NewOscillator set to a given sample rate
func NewOscillator(sampleRate int, shape Shape) (*Oscillator, error) {
	cf, ok := shapeCalcFunc[shape]
	if !ok {
		return nil, fmt.Errorf("Shape type %v not supported", shape)
	}
	return &Oscillator{
		waveMultiplier: Tau / float64(sampleRate),
		WaveFn:         cf,
	}, nil
}

func (o *Oscillator) SetFrequency(freq float64) {
	o.frequency = freq
	o.incrementer = o.waveMultiplier * o.frequency
}

func (o *Oscillator) Reset() {
	o.phase = Tau
}

func (o *Oscillator) Tick() float64 {
	tick := o.WaveFn(o.phase)
	o.phase += o.incrementer

	if o.phase >= Tau {
		o.phase -= Tau
	}

	if o.phase < 0 {
		o.phase = Tau
	}

	return tick
}

func (o *Oscillator) Samples() []byte {
	base := o.Tick()
	sampleTo16Bit := int16(base * math.MaxInt16)

	samples := []byte{
		byte(sampleTo16Bit),
		byte(sampleTo16Bit >> 8),
		byte(sampleTo16Bit),
		byte(sampleTo16Bit >> 8),
	}

	return samples
}

func sineWave(phase float64) float64 {
	return math.Sin(phase)
}

func squareWave(phase float64) float64 {
	val := -1.0
	if phase <= math.Pi {
		val = 1.0
	}

	return val
}

func triangleWave(phase float64) float64 {
	val := 2.0*(phase*(1.0/Tau)) - 1.0
	if val < 0.0 {
		val = -val
	}
	val = 2.0 * (val - 0.5)
	return val
}

func downSawWave(phase float64) float64 {
	val := 1.0 - 2.0*(phase*(1.0/Tau))
	return val
}

func upSawWave(phase float64) float64 {
	val := 2.0*(phase*(1.0/Tau)) - 1.0
	return val
}

func noiseWave(phase float64) float64 {
	return rand.Float64()*2 - 1
}
