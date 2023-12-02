package subtypes

import (
	"git.gammaspectra.live/WeebDataHoarder/swf-go/types"
	"slices"
)

type FILLSTYLEARRAY struct {
	_                      struct{} `swfFlags:"root"`
	FillStyleCount         uint8
	FillStyleCountExtended uint16      `swfCondition:"HasFillStyleCountExtended()"`
	FillStyles             []FILLSTYLE `swfCount:"FillStylesLength()"`
}

func (t *FILLSTYLEARRAY) HasFillStyleCountExtended(ctx types.ReaderContext) bool {
	return t.FillStyleCount == 0xff
}

func (t *FILLSTYLEARRAY) FillStylesLength(ctx types.ReaderContext) uint64 {
	if t.FillStyleCount == 0xff {
		return uint64(t.FillStyleCountExtended)
	}
	return uint64(t.FillStyleCount)
}

type FillStyleType uint8

const (
	FillStyleSolid = FillStyleType(0x00)

	FillStyleLinearGradient = FillStyleType(0x10)

	FillStyleRadialGradient      = FillStyleType(0x12)
	FillStyleFocalRadialGradient = FillStyleType(0x13)

	FillStyleRepeatingBitmap            = FillStyleType(0x40)
	FillStyleClippedBitmap              = FillStyleType(0x41)
	FillStyleNonSmoothedRepeatingBitmap = FillStyleType(0x42)
	FillStyleNonSmoothedClippedBitmap   = FillStyleType(0x43)
)

func (t *FillStyleType) SWFRead(r types.DataReader, ctx types.ReaderContext) (err error) {
	err = types.ReadU8(r, t)
	if err != nil {
		return err
	}
	// Bitmap smoothing only occurs in SWF version 8+.
	if ctx.Version < 8 {
		switch *t {
		case FillStyleClippedBitmap:
			*t = FillStyleNonSmoothedClippedBitmap
		case FillStyleRepeatingBitmap:
			*t = FillStyleNonSmoothedRepeatingBitmap
		}
	}
	return nil
}

type FILLSTYLE struct {
	_             struct{} `swfFlags:"root,alignend"`
	FillStyleType FillStyleType

	Color types.Color `swfCondition:"HasRGB()"`

	GradientMatrix types.MATRIX  `swfCondition:"HasGradientMatrix()"`
	Gradient       GRADIENT      `swfCondition:"HasGradient()"`
	FocalGradient  FOCALGRADIENT `swfCondition:"HasFocalGradient()"`

	BitmapId     uint16       `swfCondition:"HasBitmap()"`
	BitmapMatrix types.MATRIX `swfCondition:"HasBitmap()"`
}

func (s *FILLSTYLE) SWFDefault(ctx types.ReaderContext) {
	if slices.Contains(ctx.Flags, "Shape3") || slices.Contains(ctx.Flags, "Shape4") {
		s.Color = &types.RGBA{}
	} else {
		s.Color = &types.RGB{}
	}
}

func (s *FILLSTYLE) HasRGB(ctx types.ReaderContext) bool {
	//check first
	switch s.FillStyleType {
	case FillStyleSolid:
	case FillStyleLinearGradient:
	case FillStyleRadialGradient:
	case FillStyleFocalRadialGradient:
	case FillStyleRepeatingBitmap:
	case FillStyleClippedBitmap:
	case FillStyleNonSmoothedRepeatingBitmap:
	case FillStyleNonSmoothedClippedBitmap:
	default:
		panic("unknown fill style")

	}
	return s.FillStyleType == FillStyleSolid
}

func (s *FILLSTYLE) HasGradientMatrix(ctx types.ReaderContext) bool {
	return s.HasGradient(ctx) || s.HasFocalGradient(ctx)
}

func (s *FILLSTYLE) HasGradient(ctx types.ReaderContext) bool {
	return s.FillStyleType == FillStyleLinearGradient || s.FillStyleType == FillStyleRadialGradient
}

func (s *FILLSTYLE) HasFocalGradient(ctx types.ReaderContext) bool {
	return s.FillStyleType == FillStyleFocalRadialGradient
}

func (s *FILLSTYLE) HasBitmap(ctx types.ReaderContext) bool {
	return s.FillStyleType == FillStyleRepeatingBitmap || s.FillStyleType == FillStyleClippedBitmap || s.FillStyleType == FillStyleNonSmoothedRepeatingBitmap || s.FillStyleType == FillStyleNonSmoothedClippedBitmap
}

type LINESTYLEARRAY struct {
	_                      struct{} `swfFlags:"root"`
	LineStyleCount         uint8
	LineStyleCountExtended uint16       `swfCondition:"HasLineStyleCountExtended()"`
	LineStyles             []LINESTYLE  `swfCondition:"!HasLineStyles2()" swfCount:"LineStylesLength()"`
	LineStyles2            []LINESTYLE2 `swfCondition:"HasLineStyles2()" swfCount:"LineStylesLength()"`
}

func (t *LINESTYLEARRAY) HasLineStyleCountExtended(ctx types.ReaderContext) bool {
	return t.LineStyleCount == 0xff
}

func (t *LINESTYLEARRAY) HasLineStyles2(ctx types.ReaderContext) bool {
	return slices.Contains(ctx.Flags, "Shape4")
}

func (t *LINESTYLEARRAY) LineStylesLength(ctx types.ReaderContext) uint64 {
	if t.LineStyleCount == 0xff {
		return uint64(t.LineStyleCountExtended)
	}
	return uint64(t.LineStyleCount)
}

type LINESTYLE struct {
	Width uint16
	Color types.Color
}

func (s *LINESTYLE) SWFDefault(ctx types.ReaderContext) {
	if slices.Contains(ctx.Flags, "Shape3") {
		s.Color = &types.RGBA{}
	} else {
		s.Color = &types.RGB{}
	}
}

type LINESTYLE2 struct {
	_     struct{} `swfFlags:"root"`
	Width uint16
	Flag  struct {
		StartCapStyle      uint8 `swfBits:",2"`
		JoinStyle          uint8 `swfBits:",2"`
		HasFill            bool
		NoHScale, NoVScale bool
		PixelHinting       bool
		Reserved           uint8 `swfBits:",5"`
		NoClose            bool
		EndCapStyle        uint8 `swfBits:",2"`
	}
	MitterLimitFactor uint16     `swfCondition:"HasMitterLimitFactor()"`
	Color             types.RGBA `swfCondition:"HasColor()"`
	FillType          FILLSTYLE  `swfCondition:"Flag.HasFill"`
}

func (t *LINESTYLE2) HasMitterLimitFactor(ctx types.ReaderContext) bool {
	return t.Flag.JoinStyle == 2
}

func (t *LINESTYLE2) HasColor(ctx types.ReaderContext) bool {
	return !t.Flag.HasFill
}
