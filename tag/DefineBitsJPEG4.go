package tag

import (
	"bytes"
	"compress/zlib"
	"git.gammaspectra.live/WeebDataHoarder/swf-go/types"
	"io"
)

type DefineBitsJPEG4 struct {
	_               struct{} `swfFlags:"root"`
	CharacterId     uint16
	AlphaDataOffset uint32
	DeblockParam    types.Fixed8
	ImageData       []byte `swfCount:"AlphaDataOffset"`
	BitmapAlphaData types.UntilEndBytes
}

func (t *DefineBitsJPEG4) GetAlphaData() []byte {
	if len(t.BitmapAlphaData) == 0 {
		return nil
	}
	r, err := zlib.NewReader(bytes.NewReader(t.BitmapAlphaData))
	if err != nil {
		return nil
	}
	defer r.Close()
	buf, err := io.ReadAll(r)
	if err != nil {
		return nil
	}
	return buf
}

func (t *DefineBitsJPEG4) Code() Code {
	return RecordDefineBitsJPEG4
}
