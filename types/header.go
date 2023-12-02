package types

type HeaderSignature [3]uint8

var SignatureUncompressed = HeaderSignature{'F', 'W', 'S'}
var SignatureCompressedZLIB = HeaderSignature{'C', 'W', 'S'}
var SignatureCompressedLZMA = HeaderSignature{'Z', 'W', 'S'}

type Header struct {
	Signature  HeaderSignature
	Version    uint8
	FileLength uint32
	FrameSize  RECT
	FrameRate  Fixed8
	FrameCount uint16
}
