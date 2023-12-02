package tag

import "git.gammaspectra.live/WeebDataHoarder/swf-go/types"

type DefineBinaryData struct {
	CharacterId uint16
	Reserved    uint32
	Data        types.UntilEndBytes
}

func (t *DefineBinaryData) Code() Code {
	return RecordDefineBinaryData
}
