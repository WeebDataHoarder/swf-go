package tag

import (
	"bytes"
	"errors"
	"git.gammaspectra.live/WeebDataHoarder/swf-go/types"
	"github.com/icza/bitio"
)

type Tag interface {
	Code() Code
}

type Record struct {
	_                struct{}            `swfFlags:"root,align"`
	ctx              types.ReaderContext `swfFlags:"skip"`
	TagCodeAndLength uint16
	ExtraLength      uint32 `swfCondition:"HasExtraLength()"`
	Data             []byte `swfCount:"DataLength()"`
}

func (r *Record) SWFDefault(ctx types.ReaderContext) {
	r.ctx = ctx
}

func (r *Record) HasExtraLength(ctx types.ReaderContext) bool {
	return (r.TagCodeAndLength & 0x3f) == 0x3f
}

func (r *Record) DataLength(ctx types.ReaderContext) uint64 {
	if (r.TagCodeAndLength & 0x3f) == 0x3f {
		return uint64(r.ExtraLength)
	}
	return uint64(r.TagCodeAndLength & 0x3f)
}

func (r *Record) Decode() (readTag Tag, err error) {
	bitReader := bitio.NewReader(bytes.NewReader(r.Data))

	switch r.Code() {
	case RecordShowFrame:
		readTag = &ShowFrame{}
	case RecordPlaceObject:
		readTag = &PlaceObject{}
	case RecordRemoveObject:
		readTag = &RemoveObject{}
	case RecordPlaceObject2:
		readTag = &PlaceObject2{}
	case RecordRemoveObject2:
		readTag = &RemoveObject2{}
	case RecordPlaceObject3:
		readTag = &PlaceObject3{}
	case RecordEnd:
		readTag = &End{}
	case RecordSetBackgroundColor:
		readTag = &SetBackgroundColor{}
	case RecordProtect:
		readTag = &Protect{}
	case RecordFrameLabel:
		readTag = &FrameLabel{}
	case RecordDefineShape:
		readTag = &DefineShape{}
	case RecordDoAction:
		readTag = &DoAction{}
	case RecordDefineShape2:
		readTag = &DefineShape2{}
	case RecordDefineShape3:
		readTag = &DefineShape3{}
	case RecordDoInitAction:
		readTag = &DoInitAction{}
	case RecordFileAttributes:
		readTag = &FileAttributes{}
	case RecordMetadata:
		readTag = &Metadata{}
	case RecordDefineScalingGrid:
		readTag = &DefineScalingGrid{}
	case RecordDefineShape4:
		readTag = &DefineShape4{}
	case RecordDefineSceneAndFrameLabelData:
		readTag = &DefineSceneAndFrameLabelData{}
	case RecordDefineBits:
		readTag = &DefineBits{}
	case RecordJPEGTables:
		readTag = &JPEGTables{}
	case RecordDefineBitsJPEG2:
		readTag = &DefineBitsJPEG2{}
	case RecordDefineBitsJPEG3:
		readTag = &DefineBitsJPEG3{}
	case RecordDefineMorphShape:
		readTag = &DefineMorphShape{}
	case RecordDefineMorphShape2:
		readTag = &DefineMorphShape2{}
	case RecordDefineBitsJPEG4:
		readTag = &DefineBitsJPEG4{}
	case RecordDefineBitsLossless:
		readTag = &DefineBitsLossless{}
	case RecordDefineBitsLossless2:
		readTag = &DefineBitsLossless2{}
	case RecordDefineSprite:
		readTag = &DefineSprite{}
	case RecordDefineSound:
		readTag = &DefineSound{}
	case RecordSoundStreamHead:
		readTag = &SoundStreamHead{}
	case RecordSoundStreamHead2:
		readTag = &SoundStreamHead2{}
	case RecordSoundStreamBlock:
		readTag = &SoundStreamBlock{}
	case RecordStartSound:
		readTag = &StartSound{}
	case RecordStartSound2:
		readTag = &StartSound2{}
	case RecordDefineFont:
		readTag = &DefineFont{}
	case RecordDefineFontInfo:
		readTag = &DefineFontInfo{}
	case RecordDefineFont2:
		readTag = &DefineFont2{}
	case RecordDefineFontInfo2:
		readTag = &DefineFontInfo2{}
	case RecordDefineText:
		readTag = &DefineText{}
	case RecordDefineText2:
		readTag = &DefineText2{}
	case RecordDefineFontAlignZones:
		readTag = &DefineFontAlignZones{}
	case RecordDefineFont3:
		readTag = &DefineFont3{}
	case RecordDefineFontName:
		readTag = &DefineFontName{}
	case RecordDefineFont4:
		readTag = &DefineFont4{}
	case RecordPlaceObject4:
		readTag = &PlaceObject4{}
	case RecordDoABC:
		readTag = &DoABC{}
	case RecordEnableTelemetry:
		readTag = &EnableTelemetry{}

	case RecordExportAssets:
	case RecordImportAssets:
	case RecordImportAssets2:
	case RecordSymbolClass:

	}

	if readTag == nil {
		return nil, errors.New("could not decode tag")
	}

	err = types.ReadType(bitReader, types.ReaderContext{
		Version: r.ctx.Version,
	}, readTag)
	if err != nil {
		return nil, err
	}
	if readTag.Code() != r.Code() {
		return nil, errors.New("mismatched decoded tag code")
	}

	return readTag, nil
}

func (r *Record) Code() Code {
	return Code(r.TagCodeAndLength >> 6)
}
