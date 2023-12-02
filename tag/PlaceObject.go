package tag

import (
	"bytes"
	"git.gammaspectra.live/WeebDataHoarder/swf-go/types"
	"github.com/icza/bitio"
	"io"
)

type PlaceObject struct {
	Flag struct {
		HasColorTransform bool
	}
	CharacterId    uint16
	Depth          uint16
	Matrix         types.MATRIX
	ColorTransform *types.CXFORM
}

func (t *PlaceObject) SWFRead(r types.DataReader, ctx types.ReaderContext) (err error) {
	err = types.ReadU16(r, &t.CharacterId)
	if err != nil {
		return err
	}
	err = types.ReadU16(r, &t.Depth)
	if err != nil {
		return err
	}
	err = types.ReadType(r, ctx, &t.Matrix)
	if err != nil {
		return err
	}
	r.Align()
	data, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	if len(data) > 0 {
		ct := &types.CXFORM{}
		err = types.ReadType(bitio.NewReader(bytes.NewReader(data)), ctx, ct)
		if err != nil {
			return err
		}
		t.ColorTransform = ct
		t.Flag.HasColorTransform = true
	}
	return nil
}

func (t *PlaceObject) Code() Code {
	return RecordPlaceObject
}
