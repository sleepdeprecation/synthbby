package main

import (
	"bytes"
	"time"

	"github.com/hajimehoshi/oto/v2"
	"github.com/sleepdeprecation/synthbby/sequencer"
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

	dm := sequencer.NewDrumMachine(sampleRate, 120)
	dm.Snare.SetEnvelope(0.0, 0.15, 0, 0)
	for _, step := range dm.Snare.Steps() {
		step.Gate = sequencer.GateDiscrete
	}

	dm.Bass.SetEnvelope(0.5, 0, 1, 0.5)
	for idx, step := range dm.Bass.Steps() {
		step.Gate = sequencer.GateOpen
		if idx%2 == 0 {
			step.Gate |= sequencer.GateStart
		} else if idx%1 == 1 {
			step.Gate |= sequencer.GateEnd
		}

		switch idx / 4 {
		case 0:
			step.Pitch = 65
		case 1:
			step.Pitch = 82
		case 2:
			step.Pitch = 123
		case 3:
			step.Pitch = 98
		}
	}

	dm.Lead.SetEnvelope(0.25, 0.125, 0.7, 0.1)
	for idx, step := range dm.Lead.Steps() {
		step.Gate = sequencer.GateDiscrete
		switch idx % 4 {
		case 0:
			step.Pitch = 440
		case 1:
			step.Pitch = 523
		case 2:
			step.Pitch = 784
		case 3:
			step.Pitch = 659
		}
	}

	tmpBuf := []byte{}
	for _, sample := range dm.Render() {
		tmpBuf = append(
			tmpBuf,
			byte(sample),
			byte(sample>>8),
			byte(sample),
			byte(sample>>8),
		)
	}

	playBuf := bytes.NewReader(tmpBuf)

	player := otoCtx.NewPlayer(playBuf)
	player.Play()

	for player.IsPlaying() {
		time.Sleep(time.Millisecond)
	}

	player.Close()
}
