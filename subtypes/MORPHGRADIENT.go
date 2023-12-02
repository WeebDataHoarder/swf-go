package subtypes

import (
	"git.gammaspectra.live/WeebDataHoarder/swf-go/types"
)

type MORPHGRADIENT struct {
	_            struct{} `swfFlags:"root"`
	NumGradients uint8
	Records      []MORPHGRADRECORD `swfCount:"NumGradients"`
}

func (g MORPHGRADIENT) StartGradient() (g2 GRADIENT) {
	g2.SpreadMode = 0        //TODO
	g2.InterpolationMode = 0 //TODO
	g2.NumGradients = g.NumGradients
	for _, r := range g.Records {
		g2.Records = append(g2.Records, GRADRECORD{
			Ratio: r.StartRatio,
			Color: r.StartColor,
		})
	}
	return g2
}

func (g MORPHGRADIENT) EndGradient() (g2 GRADIENT) {
	g2.SpreadMode = 0        //TODO
	g2.InterpolationMode = 0 //TODO
	g2.NumGradients = g.NumGradients
	for _, r := range g.Records {
		g2.Records = append(g2.Records, GRADRECORD{
			Ratio: r.EndRatio,
			Color: r.EndColor,
		})
	}
	return g2
}

type MORPHGRADRECORD struct {
	StartRatio uint8
	StartColor types.RGBA
	EndRatio   uint8
	EndColor   types.RGBA
}
