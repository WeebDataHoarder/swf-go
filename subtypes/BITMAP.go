package subtypes

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"errors"
	"fmt"
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"math"
	"runtime"
	"slices"
)

type ImageBitsFormat uint8

const (
	ImageBitsFormatPaletted ImageBitsFormat = 3
	ImageBitsFormatRGB15    ImageBitsFormat = 4
	ImageBitsFormatRGB32    ImageBitsFormat = 5
)

func DecodeImageBits(data []byte, width, height int, format ImageBitsFormat, paletteSize int, hasAlpha bool) (image.Image, error) {
	r, err := zlib.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer r.Close()

	var buf []byte
	switch format {
	// 8-bit colormapped image
	case ImageBitsFormatPaletted:
		advanceWidth := make([]byte, ((uint16(width)+0b11)&(^uint16(0b11)))-uint16(width))

		if hasAlpha {
			buf = make([]byte, 4)
		} else {
			buf = make([]byte, 3)
		}
		var palette color.Palette
		for i := 0; i < paletteSize; i++ {
			_, err = io.ReadFull(r, buf[:])
			if err != nil {
				return nil, err
			}
			var a uint8 = math.MaxUint8
			if hasAlpha {
				a = buf[3]
			}
			palette = append(palette, color.RGBA{R: buf[0], G: buf[1], B: buf[2], A: a})
		}

		im := image.NewPaletted(image.Rectangle{
			Min: image.Point{},
			Max: image.Point{X: width, Y: height},
		}, palette)
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				_, err = io.ReadFull(r, buf[:1])
				if err != nil {
					return nil, err
				}
				im.SetColorIndex(x, y, buf[0])
			}

			if len(advanceWidth) > 0 {
				_, err = io.ReadFull(r, advanceWidth[:])
				if err != nil {
					return nil, err
				}
			}
		}
		return im, nil
	case ImageBitsFormatRGB15:
		if hasAlpha {
			return nil, errors.New("rgb15 not supported in alpha mode")
		}

		im := image.NewRGBA(image.Rectangle{
			Min: image.Point{},
			Max: image.Point{X: width, Y: height},
		})

		advanceWidth := make([]byte, ((uint16(width)+0b1)&(^uint16(0b1)))-uint16(width))
		buf = make([]byte, 2)

		//TODO: check if correct
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				_, err = io.ReadFull(r, buf[:])
				if err != nil {
					return nil, err
				}
				compressed := binary.BigEndian.Uint16(buf)

				im.SetRGBA(x, y, color.RGBA{
					R: uint8((((compressed>>10)&0x1f)*255 + 15) / 31),
					G: uint8((((compressed>>5)&0x1f)*255 + 15) / 31),
					B: uint8(((compressed&0x1f)*255 + 15) / 31),
					A: math.MaxUint8,
				})
			}

			if len(advanceWidth) > 0 {
				_, err = io.ReadFull(r, advanceWidth[:])
				if err != nil {
					return nil, err
				}
			}
		}
		return im, nil
	case ImageBitsFormatRGB32:
		//always read 4 bytes regardless
		buf = make([]byte, 4)

		im := image.NewRGBA(image.Rectangle{
			Min: image.Point{},
			Max: image.Point{X: width, Y: height},
		})

		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				_, err = io.ReadFull(r, buf[:])
				if err != nil {
					return nil, err
				}
				if !hasAlpha {
					buf[0] = math.MaxUint8
				}
				im.SetRGBA(x, y, color.RGBA{R: buf[1], G: buf[2], B: buf[3], A: buf[0]})
			}
		}
		return im, nil
	default:
		return nil, fmt.Errorf("unsupported lossless format %d", format)
	}
}

var bitmapHeaderJPEG = []byte{0xff, 0xd8}
var bitmapHeaderJPEGInvalid = []byte{0xff, 0xd9, 0xff, 0xd8}
var bitmapHeaderPNG = []byte{0x89, 0x50, 0x4e, 0x47, 0x0d, 0x0a, 0x1a, 0x0a}
var bitmapHeaderGIF = []byte{0x47, 0x49, 0x46, 0x38, 0x39, 0x61}

var bitmapHeaderFormats = [][]byte{
	bitmapHeaderJPEG,
	bitmapHeaderJPEGInvalid,
	bitmapHeaderPNG,
	bitmapHeaderGIF,
}

func DecodeImageBitsJPEG(imageData []byte, alphaData []byte) (image.Image, error) {
	var im image.Image
	var err error

	for i, s := range bitmapHeaderFormats {
		if bytes.Compare(s, imageData[:len(s)]) == 0 {
			if i == 0 || i == 1 {
				//jpeg
				//remove invalid data
				jpegData := removeInvalidJPEGData(imageData)
				im, _, err = image.Decode(bytes.NewReader(jpegData))
				if im != nil {
					size := im.Bounds().Size()
					if len(alphaData) == size.X*size.Y {

						newIm := image.NewRGBA(im.Bounds())
						for x := 0; x < size.X; x++ {
							for y := 0; y < size.Y; y++ {
								rI, gI, bI, _ := im.At(x, y).RGBA()

								// The JPEG data should be premultiplied alpha, but it isn't in some incorrect SWFs.
								// This means 0% alpha pixels may have color and incorrectly show as visible.
								// Flash Player clamps color to the alpha value to fix this case.
								// Only applies to DefineBitsJPEG3; DefineBitsLossless does not seem to clamp.
								a := alphaData[y*size.X+x]
								if a != 0 {
									runtime.KeepAlive(a)
								}
								r := min(uint8(rI>>8), a)
								g := min(uint8(gI>>8), a)
								b := min(uint8(bI>>8), a)
								newIm.SetRGBA(x, y, color.RGBA{
									R: r,
									G: g,
									B: b,
									A: a,
								})
							}
						}
						im = newIm
					}
				}
			} else if i == 2 {
				//png
				im, _, err = image.Decode(bytes.NewReader(imageData))
			} else if i == 3 {
				//gif
				im, _, err = image.Decode(bytes.NewReader(imageData))
			}
			break
		}
	}
	if err != nil {
		return nil, err
	}

	return im, nil
}

// removeInvalidJPEGData
// SWF19 errata p.138:
// "Before version 8 of the SWF file format, SWF files could contain an erroneous header of 0xFF, 0xD9, 0xFF, 0xD8
// before the JPEG SOI marker."
// 0xFFD9FFD8 is a JPEG EOI+SOI marker pair. Contrary to the spec, this invalid marker sequence can actually appear
// at any time before the 0xFFC0 SOF marker, not only at the beginning of the data. I believe this is a relic from
// the SWF JPEGTables tag, which stores encoding tables separately from the DefineBits image data, encased in its
// own SOI+EOI pair. When these data are glued together, an interior EOI+SOI sequence is produced. The Flash JPEG
// decoder expects this pair and ignores it, despite standard JPEG decoders stopping at the EOI.
// When DefineBitsJPEG2 etc. were introduced, the Flash encoders/decoders weren't properly adjusted, resulting in
// this sequence persisting. Also, despite what the spec says, this doesn't appear to be version checked (e.g., a
// v9 SWF can contain one of these malformed JPEGs and display correctly).
// See https://github.com/ruffle-rs/ruffle/issues/8775 for various examples.
func removeInvalidJPEGData(data []byte) (buf []byte) {
	const SOF0 uint8 = 0xC0 // Start of frame
	const RST0 uint8 = 0xD0 // Restart (we shouldn't see this before SOS, but just in case)
	const RST1 uint8 = 0xD0
	const RST2 uint8 = 0xD0
	const RST3 uint8 = 0xD0
	const RST4 uint8 = 0xD0
	const RST5 uint8 = 0xD0
	const RST6 uint8 = 0xD0
	const RST7 uint8 = 0xD7
	const SOI uint8 = 0xD8 // Start of image
	const EOI uint8 = 0xD9 // End of image

	if bytes.HasPrefix(data, bitmapHeaderJPEGInvalid) {
		data = bytes.TrimPrefix(data, bitmapHeaderJPEGInvalid)
	} else {
		// Parse the JPEG markers searching for the 0xFFD9FFD8 marker sequence to splice out.
		// We only have to search up to the SOF0 marker.
		// This might be another case where eventually we want to write our own full JPEG decoder to match Flash's decoder.
		jpegData := data
		var pos int
		for {
			if len(jpegData) < 4 {
				break
			}

			var payloadLength int

			if bytes.Compare([]byte{0xFF, EOI, 0xFF, SOI}, jpegData[:4]) == 0 {
				// Invalid EOI+SOI sequence found, splice it out.
				data = slices.Delete(slices.Clone(data), pos, pos+4)
				break
			} else if bytes.Compare([]byte{0xFF, EOI}, jpegData[:2]) == 0 { // EOI, SOI, RST markers do not include a size.

			} else if bytes.Compare([]byte{0xFF, SOI}, jpegData[:2]) == 0 {

			} else if bytes.Compare([]byte{0xFF, RST0}, jpegData[:2]) == 0 {

			} else if bytes.Compare([]byte{0xFF, RST1}, jpegData[:2]) == 0 {

			} else if bytes.Compare([]byte{0xFF, RST2}, jpegData[:2]) == 0 {

			} else if bytes.Compare([]byte{0xFF, RST3}, jpegData[:2]) == 0 {

			} else if bytes.Compare([]byte{0xFF, RST4}, jpegData[:2]) == 0 {

			} else if bytes.Compare([]byte{0xFF, RST5}, jpegData[:2]) == 0 {

			} else if bytes.Compare([]byte{0xFF, RST6}, jpegData[:2]) == 0 {

			} else if bytes.Compare([]byte{0xFF, RST7}, jpegData[:2]) == 0 {

			} else if bytes.Compare([]byte{0xFF, SOF0}, jpegData[:2]) == 0 {
				// No invalid sequence found before SOF marker, return data as-is.
				break
			} else if jpegData[0] == 0xFF {
				// Other tags include a length.
				payloadLength = int(binary.BigEndian.Uint16(jpegData[2:]))
			} else {
				// All JPEG markers should start with 0xFF.
				// So this is either not a JPEG, or we screwed up parsing the markers. Bail out.
				break
			}

			if len(jpegData) < payloadLength+2 {
				break
			}

			jpegData = jpegData[payloadLength+2:]
			pos += payloadLength + 2
		}
	}

	// Some JPEGs are missing the final EOI marker (JPEG optimizers truncate it?)
	// Flash and most image decoders will still display these images, but jpeg-decoder errors.
	// Glue on an EOI marker if its not already there and hope for the best.
	if bytes.HasSuffix(data, []byte{0xff, EOI}) {
		return data
	} else {
		//JPEG is missing EOI marker and may not decode properly
		return append(slices.Clone(data), []byte{0xff, EOI}...)
	}
}
