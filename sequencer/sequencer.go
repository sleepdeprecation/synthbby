package sequencer

import (
	"fmt"

	"github.com/sleepdeprecation/synthbby/oscillator"
)

type Sequencer struct {
	Sequence   *Sequence
	Voice      *oscillator.Voice
	SampleRate int64
	BPM        int64

	renderedSequence []byte
	position         int
}

func (s *Sequencer) Render() {
	s.position = 0

	rendered := s.Sequence.Render(s.SampleRate, s.BPM, s.Voice)
	buf := []byte{}
	for i := 0; i < len(rendered); i++ {
		buf = append(buf,
			byte(rendered[i]),
			byte(rendered[i]>>8),
			byte(rendered[i]),
			byte(rendered[i]>>8),
		)
	}

	s.renderedSequence = buf
}

func (s *Sequencer) Read(buf []byte) (int, error) {
	if len(s.renderedSequence) == 0 {
		s.Render()
	}

	tmpBuf := make([]byte, len(buf))

	for i := 0; i < len(buf); i++ {
		if len(s.renderedSequence) == 0 {
			return -1, fmt.Errorf("no rendered sequence")
		}
		pos := (s.position + i) % len(s.renderedSequence)
		tmpBuf[i] = s.renderedSequence[pos]
	}

	n := copy(buf, tmpBuf)
	s.position += n
	s.position %= len(s.renderedSequence)

	return n, nil
}

func (s *Sequencer) Close() error {
	return nil
}
