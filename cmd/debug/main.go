package main

import (
	"fmt"

	"github.com/sleepdeprecation/synthbby/synth"
)

func main() {
	env := &synth.Envelope{
		Attack:  0.5,
		Decay:   0.1,
		Sustain: 0.5,
		Release: 0.1,
	}

	samples := env.BuildStep(20, true, true)
	fmt.Println(samples)

	samples = env.BuildStep(10, true, false)
	fmt.Println(samples)

	samples = env.BuildStep(10, false, true)
	fmt.Println(samples)
}
