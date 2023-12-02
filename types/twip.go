package types

import "fmt"

const TwipFactor = 20

type Twip int64

func (t Twip) Components() (pixel int64, subPixel uint8) {
	if t < 0 {
		return int64(t / TwipFactor), uint8(-int64(t%TwipFactor) * (100 / TwipFactor))
	} else {
		return int64(t / TwipFactor), uint8(int64(t%TwipFactor) * (100 / TwipFactor))
	}
}

func (t Twip) FromFloat64(v float64) Twip {
	return Twip(v * TwipFactor)
}

func (t Twip) Float64() float64 {
	return float64(t) / TwipFactor
}

func (t Twip) Float32() float32 {
	return float32(t) / TwipFactor
}

func (t Twip) String() string {
	p, subPixel := t.Components()
	return fmt.Sprintf("%d.%02d", p, subPixel)
}
