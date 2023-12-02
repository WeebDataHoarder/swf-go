package tag

import "git.gammaspectra.live/WeebDataHoarder/swf-go/types"

type SetBackgroundColor struct {
	BackgroundColor types.RGB
}

func (s *SetBackgroundColor) Code() Code {
	return RecordSetBackgroundColor
}
