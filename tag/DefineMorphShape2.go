package tag

import (
	"git.gammaspectra.live/WeebDataHoarder/swf-go/subtypes"
	"git.gammaspectra.live/WeebDataHoarder/swf-go/types"
)

type DefineMorphShape2 struct {
	_                              struct{} `swfFlags:"root"`
	CharacterId                    uint16
	StartBounds, EndBounds         types.RECT
	StartEdgeBounds, EndEdgeBounds types.RECT
	Reserved                       uint8 `swfBits:",6"`
	UsesNonScalingStrokes          bool
	UsesScalingStrokes             bool
	Offset                         uint32
	MorphFillStyles                subtypes.MORPHFILLSTYLEARRAY
	MorphLineStyles                subtypes.MORPHLINESTYLEARRAY `swfFlags:"MorphShape2"`
	StartEdges, EndEdges           subtypes.SHAPE
}

func (t *DefineMorphShape2) Code() Code {
	return RecordDefineMorphShape2
}
