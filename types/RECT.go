package types

type RECT struct {
	_     struct{} `swfFlags:"root,alignend"`
	NBits uint8    `swfBits:",5"`
	Xmin  Twip     `swfBits:"NBits,signed"`
	Xmax  Twip     `swfBits:"NBits,signed"`
	Ymin  Twip     `swfBits:"NBits,signed"`
	Ymax  Twip     `swfBits:"NBits,signed"`
}
