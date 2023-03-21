package oscillator

type Ticker interface {
	Tick() float64
}

type ConstantTicker struct {
	Value float64
}

func (c *ConstantTicker) Tick() float64 {
	return c.Value
}
