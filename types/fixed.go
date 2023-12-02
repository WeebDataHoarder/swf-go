package types

import "math"

type Fixed16 int32

func (t Fixed16) Float64() float64 {
	return float64(t) / (math.MaxUint16 + 1)
}

func (t Fixed16) Float32() float32 {
	return float32(t) / (math.MaxUint16 + 1)
}

type Fixed8 int16

func (t Fixed8) Float64() float64 {
	return float64(t) / (math.MaxUint8 + 1)
}

func (t Fixed8) Float32() float32 {
	return float32(t) / (math.MaxUint8 + 1)
}
