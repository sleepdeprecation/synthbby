package sequencer

import "github.com/sleepdeprecation/synthbby/oscillator"

type Sequence struct {
	Notes []*Note
}

// render a sample into a 16 bit linear PCM stream
func (s *Sequence) Render(sampleRate, bpm int64, voice *oscillator.Voice) []int16 {
	samples := []int16{}

	for _, note := range s.Notes {
		voice.SetFrequency(note.Frequency)
		adsrAmp := note.Envelope.Render(sampleRate, note.Duration, bpm)
		noteBuf := make([]int16, len(adsrAmp))
		for i := 0; i < len(adsrAmp); i++ {
			noteBuf[i] = int16(float64(voice.GetSample()) * adsrAmp[i])
		}

		samples = append(samples, noteBuf...)
	}

	return samples
}
