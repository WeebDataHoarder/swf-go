package types

type RGB struct {
	Red   uint8
	Green uint8
	Blue  uint8
}

func (rgb RGB) R() uint8 {
	return rgb.Red
}

func (rgb RGB) G() uint8 {
	return rgb.Green
}

func (rgb RGB) B() uint8 {
	return rgb.Blue
}

func (rgb RGB) A() uint8 {
	return 255
}
