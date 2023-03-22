package sequencer

type StepSequencer struct {
	Steps      [16]*Step
	SampleRate int
	ClockRate  int // BPM
}

type Gate uint8

const (
	GateOff  Gate = 0
	GateOpen Gate = 1 << iota
	GateStart
	GateEnd

	GateDiscrete Gate = GateStart | GateEnd
)

func (g Gate) IsStart() bool {
	return g&GateStart != 0
}

func (g Gate) IsEnd() bool {
	return g&GateEnd != 0
}

func (g Gate) IsOff() bool {
	return g == 0
}

type Step struct {
	Pitch float64
	Gate  Gate
	On    bool
}

func NewStepSequencer(sampleRate, clockRate int) *StepSequencer {
	return &StepSequencer{
		SampleRate: sampleRate,
		ClockRate:  clockRate,
		Steps: [16]*Step{
			{}, {}, {}, {},
			{}, {}, {}, {},
			{}, {}, {}, {},
			{}, {}, {}, {},
		},
	}
}
