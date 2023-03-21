package oscillator

type PreAmpFilter interface {
	// Filter takes in one frame from an oscillator (one "tick", a value between -1 and 1), and returns a modified frame (also between -1 and 1)
	Filter(frame float64) float64
}

type PostAmpFilter interface {
	// Filter takes in one sample from a 16bit linear PCM stream and returns a modified sample
	Filter(sample int16) int16
}
