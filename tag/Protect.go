package tag

type Protect struct {
}

func (t *Protect) Code() Code {
	return RecordProtect
}
