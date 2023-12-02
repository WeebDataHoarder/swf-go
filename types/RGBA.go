package types

type RGBA struct {
	Red   uint8
	Green uint8
	Blue  uint8
	Alpha uint8
}

func (rgba RGBA) R() uint8 {
	return rgba.Red
}

func (rgba RGBA) G() uint8 {
	return rgba.Green
}

func (rgba RGBA) B() uint8 {
	return rgba.Blue
}

func (rgba RGBA) A() uint8 {
	return rgba.Alpha
}
