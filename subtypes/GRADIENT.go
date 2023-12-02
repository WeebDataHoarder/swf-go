package subtypes

import (
	"git.gammaspectra.live/WeebDataHoarder/swf-go/types"
	"slices"
)

type GradientSpreadMode uint8

const (
	GradientSpreadPad = GradientSpreadMode(iota)
	GradientSpreadReflect
	GradientSpreadRepeat
	GradientSpreadReserved
)

type GradientInterpolationMode uint8

const (
	GradientInterpolationRGB = GradientInterpolationMode(iota)
	GradientInterpolationLinearRGB
	GradientInterpolationReserved2
	GradientInterpolationReserved3
)

type GRADIENT struct {
	_                 struct{}                  `swfFlags:"root"`
	SpreadMode        GradientSpreadMode        `swfBits:",2"`
	InterpolationMode GradientInterpolationMode `swfBits:",2"`
	NumGradients      uint8                     `swfBits:",4"`
	Records           []GRADRECORD              `swfCount:"NumGradients"`

	GradientCheck struct{} `swfCondition:"GradientCheckField()"`
}

func (g *GRADIENT) GradientCheckField(ctx types.ReaderContext) bool {
	if g.NumGradients < 1 {
		panic("wrong length")
	}

	if g.SpreadMode == GradientSpreadReserved {
		// Per SWF19 p. 136, SpreadMode 3 is reserved.
		// Flash treats it as pad mode.
		g.SpreadMode = GradientSpreadPad
	}

	if g.InterpolationMode == GradientInterpolationReserved2 || g.InterpolationMode == GradientInterpolationReserved3 {
		// Per SWF19 p. 136, InterpolationMode 2 and 3 are reserved.
		// Flash treats them as normal RGB mode interpolation.
		g.InterpolationMode = GradientInterpolationRGB
	}
	return false
}

type FOCALGRADIENT struct {
	_                 struct{}                  `swfFlags:"root"`
	SpreadMode        GradientSpreadMode        `swfBits:",2"`
	InterpolationMode GradientInterpolationMode `swfBits:",2"`
	NumGradients      uint8                     `swfBits:",4"`
	Records           []GRADRECORD              `swfCount:"NumGradients"`
	FocalPoint        types.Fixed8

	GradientCheck struct{} `swfCondition:"GradientCheckField()"`
}

func (g *FOCALGRADIENT) GradientCheckField(ctx types.ReaderContext) bool {
	if g.NumGradients < 1 {
		panic("wrong length")
	}

	if g.SpreadMode == GradientSpreadReserved {
		// Per SWF19 p. 136, SpreadMode 3 is reserved.
		// Flash treats it as pad mode.
		g.SpreadMode = GradientSpreadPad
	}

	if g.InterpolationMode == GradientInterpolationReserved2 || g.InterpolationMode == GradientInterpolationReserved3 {
		// Per SWF19 p. 136, InterpolationMode 2 and 3 are reserved.
		// Flash treats them as normal RGB mode interpolation.
		g.InterpolationMode = GradientInterpolationRGB
	}
	return false
}

type GRADRECORD struct {
	Ratio uint8
	Color types.Color
}

func (g *GRADRECORD) SWFDefault(ctx types.ReaderContext) {
	if slices.Contains(ctx.Flags, "Shape3") || slices.Contains(ctx.Flags, "Shape4") {
		g.Color = &types.RGBA{}
	} else {
		g.Color = &types.RGB{}
	}
}
