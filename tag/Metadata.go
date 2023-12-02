package tag

type Metadata struct {
	Metadata string
}

func (t *Metadata) Code() Code {
	return RecordMetadata
}
