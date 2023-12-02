package tag

import (
	"git.gammaspectra.live/WeebDataHoarder/swf-go/types"
)

type DefineFontInfo2 struct {
	_           struct{} `swfFlags:"root"`
	FontId      uint16
	FontNameLen uint8
	FontName    []byte `swfCount:"FontNameLen"`
	Flag        struct {
		Reserved  uint8 `swfBits:",2"`
		SmallText bool
		ShiftJIS  bool
		ANSI      bool
		Italic    bool
		Bold      bool
		WideCodes bool
	}
	LanguageCode uint8
	CodeTable8   types.UntilEnd[uint8]  `swfCondition:"!Flag.WideCodes"`
	CodeTable16  types.UntilEnd[uint16] `swfCondition:"Flag.WideCodes"`
}

func (t *DefineFontInfo2) Code() Code {
	return RecordDefineFontInfo2
}
