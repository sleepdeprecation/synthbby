package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/sleepdeprecation/synthbby/oscillator"
	"github.com/sleepdeprecation/synthbby/sequencer"
	"github.com/sleepdeprecation/synthbby/synth"
)

const (
	screenWidth  = 640
	screenHeight = 480
	sampleRate   = 48000
)

type Game struct {
	audioContext *audio.Context
	player       *audio.Player
}

func (g *Game) Update() error {
	if g.audioContext == nil {
		g.audioContext = audio.NewContext(sampleRate)
	}
	if g.player == nil {
		// Pass the (infinite) stream to NewPlayer.
		// After calling Play, the stream never ends as long as the player object lives.
		var err error
		// osc, _ := oscillator.NewOscillator(sampleRate, oscillator.SINE)
		// env := &seq.Envelope{
		// 	Attack:  1.0,
		// 	Decay:   2.5,
		// 	Sustain: 0.75,
		// 	Release: 0.5,
		// }

		// voice := oscillator.NewSimpleVoice(sampleRate)
		// voice.PreAmpFilters = []oscillator.PreAmpFilter{
		// 	&oscillator.LowPassFilter{Cutoff: 200, SampleRate: sampleRate},
		// 	// &oscillator.HighPassFilter{Cutoff: 5000},
		// }

		// g.player, err = g.audioContext.NewPlayer(&seq.Sequencer{
		// 	SampleRate: sampleRate,
		// 	BPM:        480,
		// 	Voice:      voice,
		// 	Sequence: &seq.Sequence{
		// 		Notes: []*seq.Note{
		// 			// {Frequency: 55.0, Envelope: env, Duration: 8},
		// 			// {Frequency: 65.0, Envelope: env, Duration: 8},
		// 			// {Frequency: 98.0, Envelope: env, Duration: 8},
		// 			// {Frequency: 82.0, Envelope: env, Duration: 8},
		// 			{Frequency: float64(A3), Envelope: env, Duration: 8},
		// 			{Frequency: float64(C3), Envelope: env, Duration: 8},
		// 			{Frequency: float64(G3), Envelope: env, Duration: 8},
		// 			{Frequency: float64(E3), Envelope: env, Duration: 8},
		// 		},
		// 	},
		// })
		g.player, err = g.audioContext.NewPlayer(makeSample())
		if err != nil {
			return err
		}
		g.player.Play()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	msg := fmt.Sprintf("TPS: %0.2f\nThis is an example using infinite audio stream.", ebiten.ActualTPS())
	ebitenutil.DebugPrint(screen, msg)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Sine Wave (Ebitengine Demo)")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

func makeSample() io.Reader {
	stepSeq := sequencer.StepSequencer{
		SampleRate: sampleRate,
		ClockRate:  120,
		Steps: [16]*sequencer.Step{
			{Pitch: 440.0, Gate: sequencer.GateStart | sequencer.GateEnd},
			{Pitch: 440.0, Gate: sequencer.GateStart | sequencer.GateEnd},
			{Pitch: 220.0, Gate: sequencer.GateStart | sequencer.GateEnd},
			{Pitch: 440.0, Gate: sequencer.GateStart | sequencer.GateEnd},

			{Pitch: 440.0, Gate: sequencer.GateStart | sequencer.GateEnd},
			{Pitch: 440.0, Gate: sequencer.GateStart | sequencer.GateEnd},
			{Pitch: 220.0, Gate: sequencer.GateStart},
			{Pitch: 220.0, Gate: sequencer.GateEnd},

			{Pitch: 440.0, Gate: sequencer.GateStart | sequencer.GateEnd},
			{Pitch: 440.0, Gate: sequencer.GateStart | sequencer.GateEnd},
			{Pitch: 220.0, Gate: sequencer.GateStart},
			{Pitch: 440.0, Gate: sequencer.GateEnd},

			{Pitch: 440.0, Gate: sequencer.GateStart | sequencer.GateEnd},
			{Pitch: 440.0, Gate: sequencer.GateStart | sequencer.GateEnd},
			{Pitch: 220.0, Gate: sequencer.GateStart | sequencer.GateEnd},
			{Pitch: 440.0, Gate: sequencer.GateStart | sequencer.GateEnd},
		},
	}

	osc, err := oscillator.NewOscillator(sampleRate, oscillator.DownSaw)
	if err != nil {
		panic(err)
	}

	syn := synth.Synth{
		Envelope: &synth.Envelope{
			Attack:  0.5,
			Decay:   0.1,
			Sustain: 0.5,
			Release: 0.1,
		},
		Oscillator: osc,
	}

	tmpBuf := []byte{}
	stepFrames := stepSeq.SampleRate * (stepSeq.ClockRate / 60)
	for _, step := range stepSeq.Steps {
		samples := syn.BuildStep(
			stepFrames,
			step.Pitch,
			step.Gate.IsStart(),
			step.Gate.IsEnd(),
		)

		for sample := range samples {
			in16 := int16(sample * math.MaxInt16)
			tmpBuf = append(
				tmpBuf,
				byte(in16),
				byte(in16>>8),
				byte(in16),
				byte(in16>>8),
			)
		}
	}

	playBuf := bytes.NewReader(tmpBuf)
	return playBuf
}
