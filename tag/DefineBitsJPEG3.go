package tag

import (
	"bytes"
	"compress/zlib"
	"git.gammaspectra.live/WeebDataHoarder/swf-go/types"
	"io"
)

type DefineBitsJPEG3 struct {
	_               struct{} `swfFlags:"root"`
	CharacterId     uint16
	AlphaDataOffset uint32
	ImageData       []byte `swfCount:"AlphaDataOffset"`
	BitmapAlphaData types.UntilEndBytes
}

func (t *DefineBitsJPEG3) GetAlphaData() []byte {
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

func (t *DefineBitsJPEG3) Code() Code {
	return RecordDefineBitsJPEG3
}
