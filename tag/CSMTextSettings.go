package tag

type CSMTextSettings struct {
	TextId       uint16
	UseFlashType uint8 `swfBits:",2"`
	GridFit      uint8 `swfBits:",3"`
	Reserved     uint8 `swfBits:",3"`
	Thickness    float32
	Sharpness    float32
	Reserved2    uint8
}

func (t *CSMTextSettings) Code() Code {
	return RecordCSMTextSettings
}
