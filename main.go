package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/sleepdeprecation/synthbby/oscillator"
	seq "github.com/sleepdeprecation/synthbby/sequencer"
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
		env := &seq.Envelope{
			Attack:  0.01,
			Decay:   0.5,
			Sustain: 0.0,
			Release: 0.0,
		}

		g.player, err = g.audioContext.NewPlayer(&seq.Sequencer{
			SampleRate: sampleRate,
			BPM:        480,
			Voice:      oscillator.NewSimpleVoice(sampleRate),
			Sequence: &seq.Sequence{
				Notes: []*seq.Note{
					{Frequency: 220.0, Envelope: env, Duration: 2},
					{Frequency: 262.0, Envelope: env, Duration: 2},
					{Frequency: 392.0, Envelope: env, Duration: 2},
					{Frequency: 330.0, Envelope: env, Duration: 2},
				},
			},
		})
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
