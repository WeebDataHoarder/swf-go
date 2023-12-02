package subtypes

import (
	"git.gammaspectra.live/WeebDataHoarder/swf-go/types"
	"slices"
)

type KERNINGRECORD struct {
	KerningCodeLeft8   uint8  `swfCondition:"!Flag.WideCodes"`
	KerningCodeLeft16  uint16 `swfCondition:"Flag.WideCodes"`
	KerningCodeRight8  uint8  `swfCondition:"!Flag.WideCodes"`
	KerningCodeRight16 uint16 `swfCondition:"Flag.WideCodes"`
	KerningAdjustment  int16
}

type ZONERECORD struct {
	_           struct{} `swfFlags:"root"`
	NumZoneData uint8
	ZoneData    []ZONEDATA `swfCount:"NumZoneData"`
	Reserved    uint8      `swfBits:",6"`
	MaskY       bool
	MaskX       bool
}

type ZONEDATA struct {
	AlignmentCoordinate types.Float16
	Range               types.Float16
}

type TEXTRECORDS []TEXTRECORD

func (records *TEXTRECORDS) SWFRead(r types.DataReader, ctx types.ReaderContext) (err error) {
	glyphBits := uint8(ctx.GetNestedType("GlyphBits").Uint())
	advanceBits := uint8(ctx.GetNestedType("AdvanceBits").Uint())

	for {
		isContinue, err := types.ReadBool(r)
		if err != nil {
			return err
		}
		if !isContinue {
			break
		}
		var record TEXTRECORD
		record.Flag.Type = isContinue
		record.GlyphBits = glyphBits
		record.AdvanceBits = advanceBits
		err = types.ReadType(r, ctx, &record)
		if err != nil {
			return err
		}
		*records = append(*records, record)
		r.Align()
	}
	return nil
}

type TEXTRECORD struct {
	_    struct{} `swfFlags:"root,alignend"`
	Flag struct {
		Type       bool  `swfFlags:"skip"`
		Reserved   uint8 `swfBits:",3"`
		HasFont    bool
		HasColor   bool
		HasYOffset bool
		HasXOffset bool
	}
	FontId       uint16      `swfCondition:"Flag.HasFont"`
	Color        types.Color `swfCondition:"Flag.HasColor"`
	XOffset      int16       `swfCondition:"Flag.HasXOffset"`
	YOffset      int16       `swfCondition:"Flag.HasYOffset"`
	TextHeight   uint16      `swfCondition:"Flag.HasFont"`
	GlyphCount   uint8
	GlyphEntries []GLYPHENTRY `swfCount:"GlyphCount"`
	GlyphBits    uint8        `swfFlags:"skip"`
	AdvanceBits  uint8        `swfFlags:"skip"`
}

func (t *TEXTRECORD) SWFDefault(ctx types.ReaderContext) {
	if slices.Contains(ctx.Flags, "Text2") {
		t.Color = &types.RGBA{}
	} else {
		t.Color = &types.RGB{}
	}
}

type GLYPHENTRY struct {
	Index   uint32 `swfBits:"GlyphBits"`
	Advance int32  `swfBits:"AdvanceBits,signed"`
}
