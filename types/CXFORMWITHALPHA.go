package types

type CXFORMWITHALPHA struct {
	_    struct{} `swfFlags:"root,alignend"`
	Flag struct {
		HasAddTerms  bool
		HasMultTerms bool
	}
	NBits    uint8 `swfBits:",4"`
	Multiply struct {
		Red   Fixed8 `swfBits:"NBits,fixed"`
		Green Fixed8 `swfBits:"NBits,fixed"`
		Blue  Fixed8 `swfBits:"NBits,fixed"`
		Alpha Fixed8 `swfBits:"NBits,fixed"`
	} `swfCondition:"Flag.HasMultTerms"`
	Add struct {
		Red   int16 `swfBits:"NBits,signed"`
		Green int16 `swfBits:"NBits,signed"`
		Blue  int16 `swfBits:"NBits,signed"`
		Alpha int16 `swfBits:"NBits,signed"`
	} `swfCondition:"Flag.HasAddTerms"`
}

func (cf *CXFORMWITHALPHA) SWFDefault(ctx ReaderContext) {
	*cf = CXFORMWITHALPHA{}
	cf.Multiply.Red = 256
	cf.Multiply.Green = 256
	cf.Multiply.Blue = 256
	cf.Multiply.Alpha = 256
}
