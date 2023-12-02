package tag

type DefineFont3 struct {
	_ struct{} `swfFlags:"root"`
	// DefineFont2 Equal except that DefineFont2.Flag.WideCodes must be true
	DefineFont2
}

func (t *DefineFont3) Scale() float64 {
	return 1024 * 20
}

func (t *DefineFont3) Code() Code {
	return RecordDefineFont3
}
