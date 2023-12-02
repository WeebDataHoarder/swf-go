package tag

import (
	"git.gammaspectra.live/WeebDataHoarder/swf-go/subtypes"
	"git.gammaspectra.live/WeebDataHoarder/swf-go/types"
)

type DefineShape struct {
	_           struct{} `swfFlags:"root"`
	ShapeId     uint16
	ShapeBounds types.RECT
	Shapes      subtypes.SHAPEWITHSTYLE
}

func (t *DefineShape) Code() Code {
	return RecordDefineShape
}
