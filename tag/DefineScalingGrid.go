package tag

import "git.gammaspectra.live/WeebDataHoarder/swf-go/types"

type DefineScalingGrid struct {
	CharacterId uint16
	Splitter    types.RECT
}

func (t *DefineScalingGrid) Code() Code {
	return RecordDefineScalingGrid
}
