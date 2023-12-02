package subtypes

import (
	"git.gammaspectra.live/WeebDataHoarder/swf-go/types"
	"slices"
)

type MORPHFILLSTYLEARRAY struct {
	_                      struct{} `swfFlags:"root"`
	FillStyleCount         uint8
	FillStyleCountExtended uint16           `swfCondition:"HasFillStyleCountExtended()"`
	FillStyles             []MORPHFILLSTYLE `swfCount:"FillStylesLength()"`
}

func (t *MORPHFILLSTYLEARRAY) HasFillStyleCountExtended(ctx types.ReaderContext) bool {
	return t.FillStyleCount == 0xff
}

func (t *MORPHFILLSTYLEARRAY) FillStylesLength(ctx types.ReaderContext) uint64 {
	if t.FillStyleCount == 0xff {
		return uint64(t.FillStyleCountExtended)
	}
	return uint64(t.FillStyleCount)
}

type MORPHFILLSTYLE struct {
	_             struct{} `swfFlags:"root"`
	FillStyleType FillStyleType

	StartColor, EndColor types.RGBA `swfCondition:"HasRGB()"`

	StartGradientMatrix, EndGradientMatrix types.MATRIX  `swfCondition:"HasGradientMatrix()"`
	Gradient                               MORPHGRADIENT `swfCondition:"HasGradient()"`

	BitmapId                           uint16       `swfCondition:"HasBitmap()"`
	StartBitmapMatrix, EndBitmapMatrix types.MATRIX `swfCondition:"HasBitmap()"`
}

func (s *MORPHFILLSTYLE) HasRGB(ctx types.ReaderContext) bool {
	//check first
	switch s.FillStyleType {
	case FillStyleSolid:
	case FillStyleLinearGradient:
	case FillStyleRadialGradient:
	case FillStyleRepeatingBitmap:
	case FillStyleClippedBitmap:
	case FillStyleNonSmoothedRepeatingBitmap:
	case FillStyleNonSmoothedClippedBitmap:
	default:
		panic("unknown fill style")

	}
	return s.FillStyleType == FillStyleSolid
}

func (s *MORPHFILLSTYLE) HasGradientMatrix(ctx types.ReaderContext) bool {
	return s.HasGradient(ctx)
}

func (s *MORPHFILLSTYLE) HasGradient(ctx types.ReaderContext) bool {
	return s.FillStyleType == FillStyleLinearGradient || s.FillStyleType == FillStyleRadialGradient
}

func (s *MORPHFILLSTYLE) HasBitmap(ctx types.ReaderContext) bool {
	return s.FillStyleType == FillStyleRepeatingBitmap || s.FillStyleType == FillStyleClippedBitmap || s.FillStyleType == FillStyleNonSmoothedRepeatingBitmap || s.FillStyleType == FillStyleNonSmoothedClippedBitmap
}

type MORPHLINESTYLEARRAY struct {
	_                      struct{} `swfFlags:"root"`
	LineStyleCount         uint8
	LineStyleCountExtended uint16            `swfCondition:"HasLineStyleCountExtended()"`
	LineStyles             []MORPHLINESTYLE  `swfCondition:"!HasLineStyles2()" swfCount:"LineStylesLength()"`
	LineStyles2            []MORPHLINESTYLE2 `swfCondition:"HasLineStyles2()" swfCount:"LineStylesLength()"`
}

func (t *MORPHLINESTYLEARRAY) HasLineStyleCountExtended(ctx types.ReaderContext) bool {
	return t.LineStyleCount == 0xff
}

func (t *MORPHLINESTYLEARRAY) HasLineStyles2(ctx types.ReaderContext) bool {
	return slices.Contains(ctx.Flags, "MorphShape2")
}

func (t *MORPHLINESTYLEARRAY) LineStylesLength(ctx types.ReaderContext) uint64 {
	if t.LineStyleCount == 0xff {
		return uint64(t.LineStyleCountExtended)
	}
	return uint64(t.LineStyleCount)
}

type MORPHLINESTYLE struct {
	StartWidth uint16
	StartColor types.RGBA
	EndWidth   uint16
	EndColor   types.RGBA
}

type MORPHLINESTYLE2 struct {
	_                    struct{} `swfFlags:"root"`
	StartWidth, EndWidth uint16
	Flag                 struct {
		StartCapStyle      uint8 `swfBits:",2"`
		JoinStyle          uint8 `swfBits:",2"`
		HasFill            bool
		NoHScale, NoVScale bool
		PixelHinting       bool
		Reserved           uint8 `swfBits:",5"`
		NoClose            bool
		EndCapStyle        uint8 `swfBits:",2"`
	}
	MitterLimitFactor    uint16         `swfCondition:"HasMitterLimitFactor()"`
	StartColor, EndColor types.RGBA     `swfCondition:"!Flag.HasFill"`
	FillType             MORPHFILLSTYLE `swfCondition:"Flag.HasFill"`
}

func (t *MORPHLINESTYLE2) HasMitterLimitFactor(ctx types.ReaderContext) bool {
	return t.Flag.JoinStyle == 2
}
