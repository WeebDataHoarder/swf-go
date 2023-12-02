package tag

import (
	"git.gammaspectra.live/WeebDataHoarder/swf-go/subtypes"
	"git.gammaspectra.live/WeebDataHoarder/swf-go/types"
)

type DefineMorphShape struct {
	_                      struct{} `swfFlags:"root"`
	CharacterId            uint16
	StartBounds, EndBounds types.RECT
	Offset                 uint32
	MorphFillStyles        subtypes.MORPHFILLSTYLEARRAY
	MorphLineStyles        subtypes.MORPHLINESTYLEARRAY
	StartEdges, EndEdges   subtypes.SHAPE
}

func (t *DefineMorphShape) Code() Code {
	return RecordDefineMorphShape
}
