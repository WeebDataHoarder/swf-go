package subtypes

import "git.gammaspectra.live/WeebDataHoarder/swf-go/types"

type FILTERLIST struct {
	_               struct{} `swfFlags:"root"`
	NumberOfFilters uint8
	Filters         []FILTER `swfCount:"NumberOfFilters"`
}

type FILTER struct {
	ID   FilterId
	Data any
}

type FilterId uint8

const (
	FilterDropShadow = FilterId(iota)
	FilterBlur
	FilterGlow
	FilterBevel
	FilterGradientGlow
	FilterConvolution
	FilterColorMatrix
	FilterGradientBevel
)

func (f *FILTER) SWFRead(r types.DataReader, ctx types.ReaderContext) (err error) {
	err = types.ReadU8(r, &f.ID)
	if err != nil {
		return err
	}

	switch f.ID {
	case FilterDropShadow:
		var value DROPSHADOWFILTER
		err = types.ReadType(r, ctx, &value)
		if err != nil {
			return err
		}
		f.Data = value
	case FilterBlur:
		var value BLURFILTER
		err = types.ReadType(r, ctx, &value)
		if err != nil {
			return err
		}
		f.Data = value
	case FilterGlow:
		var value GLOWFILTER
		err = types.ReadType(r, ctx, &value)
		if err != nil {
			return err
		}
		f.Data = value
	case FilterBevel:
		var value BEVELFILTER
		err = types.ReadType(r, ctx, &value)
		if err != nil {
			return err
		}
		f.Data = value
	case FilterGradientGlow:
		var value GRADIENTGLOWFILTER
		err = types.ReadType(r, ctx, &value)
		if err != nil {
			return err
		}
		f.Data = value
	case FilterConvolution:
		var value CONVOLUTIONFILTER
		err = types.ReadType(r, ctx, &value)
		if err != nil {
			return err
		}
		f.Data = value
	case FilterColorMatrix:
		var value COLORMATRIXFILTER
		err = types.ReadType(r, ctx, &value)
		if err != nil {
			return err
		}
		f.Data = value
	case FilterGradientBevel:
		var value GRADIENTBEVELFILTER
		err = types.ReadType(r, ctx, &value)
		if err != nil {
			return err
		}
		f.Data = value

	}
	return nil
}

type DROPSHADOWFILTER struct {
	DropShadowColor types.RGBA
	BlurX, BlurY    types.Fixed16
	Angle           types.Fixed16
	Distance        types.Fixed16
	Strength        types.Fixed8
	InnerShadow     bool
	Knockout        bool
	CompositeSource bool
	Passes          uint8 `swfBits:",5"`
}

type BLURFILTER struct {
	BlurX, BlurY types.Fixed16
	Passes       uint8 `swfBits:",5"`
	Reserved     uint8 `swfBits:",3"`
}

type GLOWFILTER struct {
	GlowColor       types.RGBA
	BlurX, BlurY    types.Fixed16
	Strength        types.Fixed8
	InnerGlob       bool
	Knockout        bool
	CompositeSource bool
	Passes          uint8 `swfBits:",5"`
}

type BEVELFILTER struct {
	ShadowColor     types.RGBA
	HighLightColor  types.RGBA
	BlurX, BlurY    types.Fixed16
	Angle           types.Fixed16
	Distance        types.Fixed16
	Strength        types.Fixed8
	InnerShadow     bool
	Knockout        bool
	CompositeSource bool
	OnTop           bool
	Passes          uint8 `swfBits:",4"`
}

type GRADIENTGLOWFILTER struct {
	_               struct{} `swfFlags:"root"`
	NumColors       uint8
	GradientColors  []types.RGBA `swfCount:"NumColors"`
	GradientRatio   []uint8      `swfCount:"NumColors"`
	BlurX, BlurY    types.Fixed16
	Angle           types.Fixed16
	Distance        types.Fixed16
	Strength        types.Fixed8
	InnerShadow     bool
	Knockout        bool
	CompositeSource bool
	OnTop           bool
	Passes          uint8 `swfBits:",4"`
}

type CONVOLUTIONFILTER struct {
	_                struct{} `swfFlags:"root"`
	MatrixX, MatrixY uint8
	Divisor          float32
	Bias             float32
	Matrix           []float32 `swfCount:"MatrixSize()"`
	DefaultColor     types.RGBA
	Reserved         uint8 `swfBits:",6"`
	Clamp            bool
	PreserveAlpha    bool
}

func (f *CONVOLUTIONFILTER) MatrixSize(ctx types.ReaderContext) uint64 {
	return uint64(f.MatrixX) * uint64(f.MatrixY)
}

type COLORMATRIXFILTER struct {
	Matrix [20]float32
}

type GRADIENTBEVELFILTER struct {
	_               struct{} `swfFlags:"root"`
	NumColors       uint8
	GradientColors  []types.RGBA `swfCount:"NumColors"`
	GradientRatio   []uint8      `swfCount:"NumColors"`
	BlurX, BlurY    types.Fixed16
	Angle           types.Fixed16
	Distance        types.Fixed16
	Strength        types.Fixed8
	InnerShadow     bool
	Knockout        bool
	CompositeSource bool
	OnTop           bool
	Passes          uint8 `swfBits:",4"`
}
