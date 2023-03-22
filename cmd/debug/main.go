package main

import (
	"fmt"
)

func main() {
	// env := &synth.Envelope{
	// 	Attack:  0.5,
	// 	Decay:   0.1,
	// 	Sustain: 0.5,
	// 	Release: 0.1,
	// }

	// samples := env.BuildStep(20, true, true)
	// fmt.Println(samples)

	// samples = env.BuildStep(10, true, false)
	// fmt.Println(samples)

	// samples = env.BuildStep(10, false, true)
	// fmt.Println(samples)

	a := make([]int, 10)
	b := []int{1, 2, 3, 4, 5}

	copy(a[3:3+len(b)], b)
	fmt.Println(a)
}
