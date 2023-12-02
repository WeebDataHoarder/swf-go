package tag

import (
	"git.gammaspectra.live/WeebDataHoarder/swf-go/types"
)

type DoABC struct {
	Flags uint32
	Name  string
	Data  types.UntilEndBytes
}

func (t *DoABC) Code() Code {
	return RecordDoABC
}
