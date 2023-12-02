package tag

import (
	"git.gammaspectra.live/WeebDataHoarder/swf-go/types"
)

type DefineBits struct {
	_           struct{} `swfFlags:"root"`
	CharacterId uint16
	Data        types.UntilEndBytes
}

func (t *DefineBits) Code() Code {
	return RecordDefineBits
}
