package oscillator

import (
	"math"
)

type PreAmpFilter interface {
	// Filter takes in one frame from an oscillator (one "tick", a value between -1 and 1), and returns a modified frame (also between -1 and 1)
	Filter(frames []float64)
}

type PostAmpFilter interface {
	// Filter takes in one sample from a 16bit linear PCM stream and returns a modified sample
	Filter(samples []int16)
}

const (
	pi2 = math.Pi * 2.0
)

type LowPassFilter struct {
	Cutoff     float64 // frequency
	SampleRate int64

	inited bool
	dt     float64
}

func (f *LowPassFilter) init() {
	f.dt = 1.0 / float64(f.SampleRate)

	f.inited = true
}

func (f *LowPassFilter) Filter(frame float64) float64 {
	if !f.inited {
		f.init()
	}

	numerator := pi2 * f.Cutoff
	denominator := frame + (numerator)

	return numerator / denominator
}

type HighPassFilter struct {
	Cutoff float64 // frequency
}

func (f *HighPassFilter) Filter(frame float64) float64 {
	denominator := pi2 * f.Cutoff
	numerator := frame + (denominator)

	return numerator / denominator
}
