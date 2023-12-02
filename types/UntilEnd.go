package types

import (
	"errors"
	"io"
)

type UntilEndBytes = UntilEnd[byte]

type UntilEnd[T any] []T

func (b *UntilEnd[T]) SWFRead(r DataReader, ctx ReaderContext) (err error) {
	for {
		var data T
		err := ReadType(r, ctx, &data)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}
		*b = append(*b, data)
	}

	return nil
}
