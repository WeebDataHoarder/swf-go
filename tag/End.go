package tag

type End struct {
}

func (t *End) Code() Code {
	return RecordEnd
}
