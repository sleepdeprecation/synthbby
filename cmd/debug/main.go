package main

import (
	"fmt"

	"github.com/sleepdeprecation/synthbby/sequencer"
)

func main() {
	env := &sequencer.Envelope{
		Attack:  1.0,
		Decay:   1.0,
		Sustain: 0.8,
		Release: 0.8,
	}

	samples := env.Render(10, 4, 60)
	fmt.Println(samples)
}
