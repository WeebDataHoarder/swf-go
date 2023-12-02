package types

type MATRIX struct {
	_                        struct{} `swfFlags:"root,alignend"`
	HasScale                 bool
	NScaleBits               uint8   `swfCondition:"HasScale" swfBits:",5"`
	ScaleX, ScaleY           Fixed16 `swfCondition:"HasScale" swfBits:"NScaleBits,fixed"`
	HasRotate                bool
	NRotateBits              uint8   `swfCondition:"HasRotate" swfBits:",5"`
	RotateSkew0, RotateSkew1 Fixed16 `swfCondition:"HasRotate" swfBits:"NRotateBits,fixed"`
	NTranslateBits           uint8   `swfBits:",5"`
	TranslateX, TranslateY   Twip    `swfBits:"NTranslateBits,signed"`
}

func (matrix *MATRIX) SWFDefault(ctx ReaderContext) {
	*matrix = MATRIX{}
	matrix.ScaleX = 1 << 16
	matrix.ScaleY = 1 << 16
}
