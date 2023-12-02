package tag

import (
	"git.gammaspectra.live/WeebDataHoarder/swf-go/subtypes"
	"git.gammaspectra.live/WeebDataHoarder/swf-go/types"
)

type DefineFont2 struct {
	_      struct{} `swfFlags:"root"`
	FontId uint16
	Flag   struct {
		HasLayout   bool
		ShiftJIS    bool
		SmallText   bool
		ANSI        bool
		WideOffsets bool
		WideCodes   bool
		Italic      bool
		Bold        bool
	}
	LanguageCode uint8
	FontNameLen  uint8
	FontName     []byte `swfCount:"FontNameLen"`
	NumGlyphs    uint16

	OffsetTable16     []uint16 `swfCount:"NumGlyphs" swfCondition:"!Flag.WideOffsets"`
	CodeTableOffset16 uint16   `swfCondition:"!Flag.WideOffsets"`

	OffsetTable32     []uint32 `swfCount:"NumGlyphs" swfCondition:"Flag.WideOffsets"`
	CodeTableOffset32 uint32   `swfCondition:"Flag.WideOffsets"`

	ShapeTable []subtypes.SHAPE `swfCount:"NumGlyphs"`

	CodeTable8  []uint8  `swfCount:"NumGlyphs" swfCondition:"!Flag.WideCodes"`
	CodeTable16 []uint16 `swfCount:"NumGlyphs" swfCondition:"Flag.WideCodes"`

	FontAscent       uint16                   `swfCondition:"Flag.HasLayout"`
	FontDescent      uint16                   `swfCondition:"Flag.HasLayout"`
	FontLeading      int16                    `swfCondition:"Flag.HasLayout"`
	FontAdvanceTable []int16                  `swfCount:"NumGlyphs" swfCondition:"Flag.HasLayout"`
	FontBoundsTable  []types.RECT             `swfCount:"NumGlyphs" swfCondition:"Flag.HasLayout"`
	KerningCount     uint16                   `swfCondition:"Flag.HasLayout"`
	KerningTable     []subtypes.KERNINGRECORD `swfCount:"KerningCount" swfCondition:"Flag.HasLayout"`
}

func (t *DefineFont2) Scale() float64 {
	return 1024
}

func (t *DefineFont2) Code() Code {
	return RecordDefineFont2
}
