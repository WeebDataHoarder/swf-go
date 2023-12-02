package tag

import (
	"git.gammaspectra.live/WeebDataHoarder/swf-go/types"
)

type DefineFont4 struct {
	_      struct{} `swfFlags:"root"`
	FontId uint16
	Flag   struct {
		Reserved    uint8 `swfBits:",5"`
		HasFontData bool
		Italic      bool
		Bold        bool
	}
	Name     string
	FontData types.UntilEndBytes `swfCondition:"Flag.HasFontData"`
}

func (t *DefineFont4) Scale() float64 {
	return 1024 * 20
}

func (t *DefineFont4) Code() Code {
	return RecordDefineFont4
}
