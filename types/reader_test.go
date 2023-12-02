package types

import (
	"bytes"
	"github.com/icza/bitio"
	"testing"
)

func TestReadSBMinus1(t *testing.T) {
	val := int8(-1)
	data := []byte{uint8(val)}
	r := bitio.NewReader(bytes.NewReader(data))

	result, err := ReadSB[int64](r, 8)
	if err != nil {
		t.Fatal(err)
	}
	if result != int64(val) {
		t.Fatal("does not match")
	}
}

func TestReadSBMinus7(t *testing.T) {
	val := int8(-7)
	data := []byte{uint8(val)}
	r := bitio.NewReader(bytes.NewReader(data))

	result, err := ReadSB[int64](r, 8)
	if err != nil {
		t.Fatal(err)
	}
	if result != int64(val) {
		t.Fatal("does not match")
	}
}

func TestReadSB127(t *testing.T) {
	val := int8(127)
	data := []byte{uint8(val)}
	r := bitio.NewReader(bytes.NewReader(data))

	result, err := ReadSB[int64](r, 8)
	if err != nil {
		t.Fatal(err)
	}
	if result != int64(val) {
		t.Fatal("does not match")
	}
}

func TestReadSBMinus128(t *testing.T) {
	val := int8(-128)
	data := []byte{uint8(val)}
	r := bitio.NewReader(bytes.NewReader(data))

	result, err := ReadSB[int64](r, 8)
	if err != nil {
		t.Fatal(err)
	}
	if result != int64(val) {
		t.Fatal("does not match")
	}
}
