package main

import (
	"bytes"
	"time"

	"github.com/hajimehoshi/oto/v2"
	"github.com/sleepdeprecation/synthbby/oscillator"
	"github.com/sleepdeprecation/synthbby/sequencer"
	"github.com/sleepdeprecation/synthbby/synth"
)

func main() {
	sampleRate := 44100
	channels := 2
	bitDepth := 2 // 16 bit audio

	otoCtx, readyChan, err := oto.NewContext(sampleRate, channels, bitDepth)
	if err != nil {
		panic(err)
	}

	<-readyChan

	stepSeq := sequencer.StepSequencer{
		SampleRate: sampleRate,
		ClockRate:  120,
		Steps: [16]*sequencer.Step{
			{Pitch: 219.0, Gate: sequencer.GateStart | sequencer.GateEnd},
			{Pitch: 219.0, Gate: sequencer.GateStart | sequencer.GateEnd},
			{Pitch: 320.0, Gate: sequencer.GateStart | sequencer.GateEnd},
			{Pitch: 219.0, Gate: sequencer.GateStart | sequencer.GateEnd},

			{Pitch: 219.0, Gate: sequencer.GateStart | sequencer.GateEnd},
			{Pitch: 219.0, Gate: sequencer.GateStart | sequencer.GateEnd},
			{Pitch: 330.0, Gate: sequencer.GateStart},
			{Pitch: 330.0, Gate: sequencer.GateEnd},

			{Pitch: 219.0, Gate: sequencer.GateStart | sequencer.GateEnd},
			// {Pitch: 219.0, Gate: sequencer.GateStart | sequencer.GateEnd},
			{Pitch: 330.0, Gate: sequencer.GateStart},
			{Pitch: 275.0, Gate: 0},
			{Pitch: 219.0, Gate: sequencer.GateEnd},

			{Pitch: 219.0, Gate: sequencer.GateStart | sequencer.GateEnd},
			{Pitch: 219.0, Gate: sequencer.GateStart | sequencer.GateEnd},
			{Pitch: 330.0, Gate: sequencer.GateStart | sequencer.GateEnd},
			{Pitch: 219.0, Gate: sequencer.GateStart | sequencer.GateEnd},
		},
	}

	osc, err := oscillator.NewOscillator(sampleRate, oscillator.Sine)
	if err != nil {
		panic(err)
	}

	syn := synth.Synth{
		Envelope: &synth.Envelope{
			Attack:  0.5,
			Decay:   0.2,
			Sustain: 0.2,
			Release: 0.1,
		},
		Oscillator: osc,
	}

	tmpBuf := []byte{}
	stepFrames := stepSeq.SampleRate / (stepSeq.ClockRate / 60)
	for _, step := range stepSeq.Steps {
		samples := syn.BuildStep(
			stepFrames,
			step.Pitch,
			step.Gate.IsStart(),
			step.Gate.IsEnd(),
		)

		for _, sample := range samples {
			tmpBuf = append(
				tmpBuf,
				byte(sample),
				byte(sample>>8),
				byte(sample),
				byte(sample>>8),
			)
		}
	}

	playBuf := bytes.NewReader(tmpBuf)

	player := otoCtx.NewPlayer(playBuf)
	player.Play()

	for player.IsPlaying() {
		time.Sleep(time.Millisecond)
	}

	player.Close()
}
