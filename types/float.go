package types

import (
	"github.com/x448/float16"
)

// Float16 TODO: check if proper values
type Float16 float16.Float16

func (f *Float16) SWFRead(r DataReader, ctx ReaderContext) (err error) {
	err = ReadU16(r, f)
	if err != nil {
		return err
	}
	return nil
}
