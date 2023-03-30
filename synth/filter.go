package synth

import (
	"math"
)

type Filter func([]float64)

func BitCrusher(numBits int) Filter {
	nBitsF := float64(numBits)
	return func(samples []float64) {
		for idx, sample := range samples {
			samples[idx] = math.Round(sample*nBitsF) / nBitsF
		}
	}
}

func Mute(samples []float64) {
	for idx := range samples {
		samples[idx] = 0
	}
}

func Overdrive(factor float64) Filter {
	return func(samples []float64) {
		for idx, sample := range samples {
			samples[idx] = sample * factor
		}
	}
}

func WaveFolder(factor float64) Filter {
	return func(samples []float64) {
		for idx, sample := range samples {
			if sample > factor {
				diff := sample - factor
				samples[idx] = factor - diff
			} else if sample < -factor {
				diff := factor - (sample * -1)
				samples[idx] = -factor + diff
			}
		}
	}
}

func WaveFoldBacker(top, bottom float64) Filter {
	return func(samples []float64) {

		for idx, sample := range samples {
			positive := sample > 0
			diff := 0.0
			current := math.Abs(sample)

			for current > top || current < bottom {
				if current > top {
					diff = current - top
					current = current - diff
				} else if current < bottom {
					diff = bottom - current
					current = current + diff
				}
			}

			if !positive {
				current = current * -1
			}

			samples[idx] = current
		}
	}
}
