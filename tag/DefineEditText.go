package tag

import (
	"git.gammaspectra.live/WeebDataHoarder/swf-go/types"
	"image/color"
)

type DefineEditText struct {
	_           struct{} `swfFlags:"root"`
	CharacterId uint16
	Bounds      types.RECT
	Flag        struct {
		HasText      bool
		WordWrap     bool
		Multiline    bool
		Password     bool
		ReadOnly     bool
		HasTextColor bool
		HasMaxLength bool
		HasFont      bool
		HasFontClass bool
		AutoSize     bool
		HasLayout    bool
		NoSelect     bool
		Border       bool
		WasStatic    bool
		HTML         bool
		UseOutlines  bool
	}

	FontId       uint16     `swfCondition:"Flag.HasFont"`
	FontClass    string     `swfCondition:"Flag.HasFontClass"`
	FontHeight   uint16     `swfCondition:"Flag.HasFont"`
	TextColor    color.RGBA `swfCondition:"Flag.HasTextColor"`
	MaxLength    uint16     `swfCondition:"Flag.HasMaxLength"`
	Align        uint8      `swfCondition:"Flag.HasLayout"`
	LeftMargin   uint16     `swfCondition:"Flag.HasLayout"`
	RightMargin  uint16     `swfCondition:"Flag.HasLayout"`
	Indent       uint16     `swfCondition:"Flag.HasLayout"`
	Leading      int16      `swfCondition:"Flag.HasLayout"`
	VariableName string
	InitialText  string `swfCondition:"Flag.HasText"`
}

func (t *DefineEditText) Code() Code {
	return RecordDefineEditText
}
