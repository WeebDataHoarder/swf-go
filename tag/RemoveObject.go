package tag

type RemoveObject struct {
	CharacterId uint16
	Depth       uint16
}

func (t *RemoveObject) Code() Code {
	return RecordRemoveObject
}
