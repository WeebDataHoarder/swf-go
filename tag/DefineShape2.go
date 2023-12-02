package tag

import (
	"git.gammaspectra.live/WeebDataHoarder/swf-go/subtypes"
	"git.gammaspectra.live/WeebDataHoarder/swf-go/types"
)

type DefineShape2 struct {
	_           struct{} `swfFlags:"root,align"`
	ShapeId     uint16
	ShapeBounds types.RECT
	Shapes      subtypes.SHAPEWITHSTYLE
}

func (t *DefineShape2) Code() Code {
	return RecordDefineShape2
}
