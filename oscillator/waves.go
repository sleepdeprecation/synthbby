package oscillator

import (
	"math"
	"math/rand"
)

func SineWave(phase float64) float64 {
	return math.Sin(phase)
}

func SquareWave(phase float64) float64 {
	val := -1.0
	if phase <= math.Pi {
		val = 1.0
	}

	return val
}

func TriangleWave(phase float64) float64 {
	val := 2.0*(phase*(1.0/Tau)) - 1.0
	if val < 0.0 {
		val = -val
	}
	val = 2.0 * (val - 0.5)
	return val
}

func DownSawWave(phase float64) float64 {
	val := 1.0 - 2.0*(phase*(1.0/Tau))
	return val
}

func UpSawWave(phase float64) float64 {
	val := 2.0*(phase*(1.0/Tau)) - 1.0
	return val
}

func NoiseWave(phase float64) float64 {
	return rand.Float64()*2 - 1
}
