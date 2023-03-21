package sequencer

type StepSequencer struct {
	Steps      [16]*Step
	SampleRate int
	ClockRate  int // BPM
}

type Gate uint8

const (
	GateStart Gate = 1 << iota
	GateEnd
)

func (g Gate) IsStart() bool {
	return g&GateStart != 0
}

func (g Gate) IsEnd() bool {
	return g&GateEnd != 0
}

type Step struct {
	Pitch float64
	Gate  Gate
}
