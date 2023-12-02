package tag

type RemoveObject2 struct {
	Depth uint16
}

func (t *RemoveObject2) Code() Code {
	return RecordRemoveObject2
}
