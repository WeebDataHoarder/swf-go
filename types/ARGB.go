package types

type ARGB struct {
	Alpha uint8
	Red   uint8
	Green uint8
	Blue  uint8
}

func (argb ARGB) R() uint8 {
	return argb.Red
}

func (argb ARGB) G() uint8 {
	return argb.Green
}

func (argb ARGB) B() uint8 {
	return argb.Blue
}

func (argb ARGB) A() uint8 {
	return argb.Alpha
}
