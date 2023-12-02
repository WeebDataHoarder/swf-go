package subtypes

import (
	"git.gammaspectra.live/WeebDataHoarder/swf-go/types"
	"reflect"
)

type SHAPE struct {
	_        struct{} `swfFlags:"root"`
	FillBits uint8    `swfBits:",4"`
	LineBits uint8    `swfBits:",4"`
	Records  SHAPERECORDS
}

type SHAPEWITHSTYLE struct {
	_          struct{} `swfFlags:"root"`
	FillStyles FILLSTYLEARRAY
	LineStyles LINESTYLEARRAY
	FillBits   uint8 `swfBits:",4"`
	LineBits   uint8 `swfBits:",4"`
	Records    SHAPERECORDS
}

type SHAPERECORDS []SHAPERECORD

func (records *SHAPERECORDS) SWFRead(r types.DataReader, ctx types.ReaderContext) (err error) {
	fillBits := uint8(ctx.GetNestedType("FillBits").Uint())
	lineBits := uint8(ctx.GetNestedType("LineBits").Uint())
	*records = make(SHAPERECORDS, 0, 512)
	for {
		isEdge, err := types.ReadBool(r)
		if err != nil {
			return err
		}

		if !isEdge {
			rec := StyleChangeRecord{}

			rec.FillBits = fillBits
			rec.LineBits = lineBits

			err = rec.SWFRead(r, types.ReaderContext{
				Version: ctx.Version,
				Root:    reflect.ValueOf(rec),
				Flags:   ctx.Flags,
			})
			if err != nil {
				return err
			}

			if rec.IsEnd() {
				//end record
				*records = append(*records, &EndShapeRecord{})
				break
			}

			//store new value
			fillBits = rec.FillBits
			lineBits = rec.LineBits

			*records = append(*records, &rec)
		} else {
			isStraight, err := types.ReadBool(r)
			if err != nil {
				return err
			}
			if isStraight {
				rec := StraightEdgeRecord{}
				err = rec.SWFRead(r, types.ReaderContext{
					Version: ctx.Version,
					Root:    reflect.ValueOf(rec),
					Flags:   ctx.Flags,
				})
				if err != nil {
					return err
				}
				*records = append(*records, &rec)
			} else {
				rec := CurvedEdgeRecord{}
				err = rec.SWFRead(r, types.ReaderContext{
					Version: ctx.Version,
					Root:    reflect.ValueOf(rec),
					Flags:   ctx.Flags,
				})
				if err != nil {
					return err
				}
				*records = append(*records, &rec)
			}
		}
	}

	r.Align()

	return nil
}

type EndShapeRecord struct {
}

func (s *EndShapeRecord) RecordType() RecordType {
	return RecordTypeEndShape
}

type StyleChangeRecord struct {
	_    struct{} `swfFlags:"root"`
	Flag struct {
		NewStyles  bool
		LineStyle  bool
		FillStyle1 bool
		FillStyle0 bool
		MoveTo     bool
	}

	MoveBits               uint8          `swfBits:",5" swfCondition:"Flag.MoveTo"`
	MoveDeltaX, MoveDeltaY types.Twip     `swfBits:"MoveBits,signed" swfCondition:"Flag.MoveTo"`
	FillStyle0             uint16         `swfBits:"FillBits" swfCondition:"Flag.FillStyle0"`
	FillStyle1             uint16         `swfBits:"FillBits" swfCondition:"Flag.FillStyle1"`
	LineStyle              uint16         `swfBits:"LineBits" swfCondition:"Flag.LineStyle"`
	FillStyles             FILLSTYLEARRAY `swfFlags:"align" swfCondition:"Flag.NewStyles"`
	LineStyles             LINESTYLEARRAY `swfCondition:"Flag.NewStyles"`

	FillBits uint8 `swfBits:",4" swfCondition:"Flag.NewStyles"`
	LineBits uint8 `swfBits:",4" swfCondition:"Flag.NewStyles"`
}

func (rec *StyleChangeRecord) IsEnd() bool {
	return !rec.Flag.NewStyles && !rec.Flag.LineStyle && !rec.Flag.FillStyle1 && !rec.Flag.FillStyle0 && !rec.Flag.MoveTo
}

func (rec *StyleChangeRecord) SWFRead(r types.DataReader, ctx types.ReaderContext) (err error) {
	//this is not necessary, but speedups read massively
	rec.Flag.NewStyles, err = types.ReadBool(r)
	if err != nil {
		return err
	}
	rec.Flag.LineStyle, err = types.ReadBool(r)
	if err != nil {
		return err
	}
	rec.Flag.FillStyle1, err = types.ReadBool(r)
	if err != nil {
		return err
	}
	rec.Flag.FillStyle0, err = types.ReadBool(r)
	if err != nil {
		return err
	}
	rec.Flag.MoveTo, err = types.ReadBool(r)
	if err != nil {
		return err
	}

	if rec.Flag.MoveTo {
		rec.MoveBits, err = types.ReadUB[uint8](r, 5)
		if err != nil {
			return err
		}
		rec.MoveDeltaX, err = types.ReadSB[types.Twip](r, uint64(rec.MoveBits))
		if err != nil {
			return err
		}
		rec.MoveDeltaY, err = types.ReadSB[types.Twip](r, uint64(rec.MoveBits))
		if err != nil {
			return err
		}
	}

	if rec.Flag.FillStyle0 {
		rec.FillStyle0, err = types.ReadUB[uint16](r, uint64(rec.FillBits))
		if err != nil {
			return err
		}
	}

	if rec.Flag.FillStyle1 {
		rec.FillStyle1, err = types.ReadUB[uint16](r, uint64(rec.FillBits))
		if err != nil {
			return err
		}
	}

	if rec.Flag.LineStyle {
		rec.LineStyle, err = types.ReadUB[uint16](r, uint64(rec.LineBits))
		if err != nil {
			return err
		}
	}

	if rec.Flag.NewStyles {
		r.Align()

		err = types.ReadType(r, types.ReaderContext{
			Version: ctx.Version,
			Root:    reflect.ValueOf(rec.FillStyles),
			Flags:   ctx.Flags,
		}, &rec.FillStyles)
		if err != nil {
			return err
		}
		err = types.ReadType(r, types.ReaderContext{
			Version: ctx.Version,
			Root:    reflect.ValueOf(rec.LineStyles),
			Flags:   ctx.Flags,
		}, &rec.LineStyles)
		if err != nil {
			return err
		}

		rec.FillBits, err = types.ReadUB[uint8](r, 4)
		if err != nil {
			return err
		}
		rec.LineBits, err = types.ReadUB[uint8](r, 4)
		if err != nil {
			return err
		}
	}

	return nil
}

func (rec *StyleChangeRecord) RecordType() RecordType {
	return RecordTypeStyleChange
}

type StraightEdgeRecord struct {
	_ struct{} `swfFlags:"root"`

	NumBits uint8 `swfBits:",4"`

	GeneralLine bool
	VertLine    bool `swfCondition:"HasVertLine()"`

	DeltaX types.Twip `swfBits:"NumBits+2,signed" swfCondition:"HasDeltaX()"`
	DeltaY types.Twip `swfBits:"NumBits+2,signed" swfCondition:"HasDeltaY()"`
}

func (s *StraightEdgeRecord) SWFRead(r types.DataReader, ctx types.ReaderContext) (err error) {
	//this is not necessary, but speedups read massively
	s.NumBits, err = types.ReadUB[uint8](r, 4)
	if err != nil {
		return err
	}
	s.GeneralLine, err = types.ReadBool(r)
	if err != nil {
		return err
	}
	if s.HasVertLine(ctx) {
		s.VertLine, err = types.ReadBool(r)
		if err != nil {
			return err
		}
	}
	if s.HasDeltaX(ctx) {
		s.DeltaX, err = types.ReadSB[types.Twip](r, uint64(s.NumBits+2))
		if err != nil {
			return err
		}
	}
	if s.HasDeltaY(ctx) {
		s.DeltaY, err = types.ReadSB[types.Twip](r, uint64(s.NumBits+2))
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *StraightEdgeRecord) HasVertLine(ctx types.ReaderContext) bool {
	return !s.GeneralLine
}

func (s *StraightEdgeRecord) HasDeltaX(ctx types.ReaderContext) bool {
	return s.GeneralLine || !s.VertLine
}

func (s *StraightEdgeRecord) HasDeltaY(ctx types.ReaderContext) bool {
	return s.GeneralLine || s.VertLine
}

func (s *StraightEdgeRecord) RecordType() RecordType {
	return RecordTypeStraightEdge
}

type CurvedEdgeRecord struct {
	_ struct{} `swfFlags:"root"`

	NumBits uint8 `swfBits:",4"`

	ControlDeltaX, ControlDeltaY types.Twip `swfBits:"NumBits+2,signed"`
	AnchorDeltaX, AnchorDeltaY   types.Twip `swfBits:"NumBits+2,signed"`
}

func (s *CurvedEdgeRecord) SWFRead(r types.DataReader, ctx types.ReaderContext) (err error) {
	//this is not necessary, but speedups read massively
	s.NumBits, err = types.ReadUB[uint8](r, 4)
	if err != nil {
		return err
	}
	s.ControlDeltaX, err = types.ReadSB[types.Twip](r, uint64(s.NumBits+2))
	if err != nil {
		return err
	}
	s.ControlDeltaY, err = types.ReadSB[types.Twip](r, uint64(s.NumBits+2))
	if err != nil {
		return err
	}
	s.AnchorDeltaX, err = types.ReadSB[types.Twip](r, uint64(s.NumBits+2))
	if err != nil {
		return err
	}
	s.AnchorDeltaY, err = types.ReadSB[types.Twip](r, uint64(s.NumBits+2))
	if err != nil {
		return err
	}
	return nil
}

func (s *CurvedEdgeRecord) RecordType() RecordType {
	return RecordTypeCurvedEdge
}

type RecordType uint8

const (
	RecordTypeEndShape = RecordType(iota)
	RecordTypeStyleChange
	RecordTypeStraightEdge
	RecordTypeCurvedEdge
)

type SHAPERECORD interface {
	RecordType() RecordType
}
