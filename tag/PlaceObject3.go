package tag

import (
	"git.gammaspectra.live/WeebDataHoarder/swf-go/subtypes"
	"git.gammaspectra.live/WeebDataHoarder/swf-go/types"
)

type BlendMode uint8

const (
	BlendNormal0 = BlendMode(iota)
	BlendNormal1
	BlendLayer
	BlendMultiply
	BlendScreen
	BlendLighten
	BlendDarken
	BlendDifference
	BlendAdd
	BlendSubtract
	BlendInvert
	BlendAlpha
	BlendErase
	BlendOverlay
	BlenHardlight
)

type PlaceObject3 struct {
	_    struct{} `swfFlags:"root,align"`
	Flag struct {
		HasClipActions    bool
		HasClipDepth      bool
		HasName           bool
		HasRatio          bool
		HasColorTransform bool
		HasMatrix         bool
		HasCharacter      bool
		Move              bool
		Reserved          bool
		OpaqueBackground  bool
		HasVisible        bool
		HasImage          bool
		HasClassName      bool
		HasCacheAsBitmap  bool
		HasBlendMode      bool
		HasFilterList     bool
	}
	Depth             uint16
	ClassName         string                `swfCondition:"HasClassName()"`
	CharacterId       uint16                `swfCondition:"Flag.HasCharacter"`
	Matrix            types.MATRIX          `swfCondition:"Flag.HasMatrix"`
	ColorTransform    types.CXFORMWITHALPHA `swfCondition:"Flag.HasColorTransform"`
	Ratio             uint16                `swfCondition:"Flag.HasRatio"`
	Name              string                `swfCondition:"Flag.HasName"`
	ClipDepth         uint16                `swfCondition:"Flag.HasClipDepth"`
	SurfaceFilterList subtypes.FILTERLIST   `swfCondition:"Flag.HasFilterList"`
	BlendMode         BlendMode             `swfCondition:"Flag.HasBlendMode"`
	BitmapCache       uint8                 `swfCondition:"Flag.HasCacheAsBitmap"`
	Visible           uint8                 `swfCondition:"Flag.HasVisible"`
	BackgroundColor   types.RGBA            `swfCondition:"Flag.OpaqueBackground"`
	ClipActions       subtypes.CLIPACTIONS  `swfCondition:"Flag.HasClipActions"`
}

func (t *PlaceObject3) HasClassName(ctx types.ReaderContext) bool {
	return t.Flag.HasClassName || (t.Flag.HasName && t.Flag.HasImage)
}

func (t *PlaceObject3) Code() Code {
	return RecordPlaceObject3
}
