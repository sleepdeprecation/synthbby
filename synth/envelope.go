package synth

type Envelope struct {
	// Attack, Decay, and Release are percents of a step
	Attack  float64
	Decay   float64
	Release float64

	// Sustain is an amplitude level, and must be between 0 and 1
	Sustain float64
}

func (e *Envelope) BuildStep(numFrames int, gateOff, gateOpen, gateClose bool) []float64 {
	frames := make([]float64, numFrames)

	if gateOff {
		return frames
	}

	attackLength := 0
	decayStart := 0
	decayEnd := 0
	decayLength := 0
	releaseStart := numFrames
	releaseLength := 0
	if gateOpen {
		attackLength = int(e.Attack * float64(numFrames))
		decayLength = int(e.Decay * float64(numFrames))
		decayStart = attackLength
		decayEnd = attackLength + decayLength
	}

	if gateClose {
		releaseLength = int(e.Release * float64(numFrames))
		if (decayEnd + releaseLength) > numFrames {
			releaseStart = decayEnd
		} else {
			releaseStart = numFrames - releaseLength
		}
	}

	for pos := 0; pos < numFrames; pos++ {
		switch {
		case pos < attackLength:
			frames[pos] = float64(pos) / float64(attackLength)
		case pos < decayEnd:
			decayPos := pos - decayStart
			frames[pos] = 1 - ((float64(decayPos) / float64(decayLength)) * (1 - e.Sustain))
		case pos >= releaseStart:
			releasePos := pos - releaseStart
			frames[pos] = (1 - (float64(releasePos) / float64(releaseLength))) * e.Sustain
		default:
			frames[pos] = e.Sustain
		}
	}

	return frames
}
