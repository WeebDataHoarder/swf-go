package tag

import (
	"git.gammaspectra.live/WeebDataHoarder/swf-go/types"
	"io"
)

type FrameLabel struct {
	Name      string
	HasAnchor bool
	Anchor    uint8
}

func (t *FrameLabel) SWFRead(r types.DataReader, ctx types.ReaderContext) (err error) {
	t.Name, err = types.ReadNullTerminatedString(r, ctx.Version)
	if err != nil {
		return err
	}
	r.Align()
	data, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	if len(data) > 0 {
		t.Anchor = data[0]
		t.HasAnchor = true
	}
	return nil
}

func (t *FrameLabel) Code() Code {
	return RecordFrameLabel
}
