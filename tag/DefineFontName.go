package tag

type DefineFontName struct {
	_         struct{} `swfFlags:"root"`
	FontId    uint16
	Name      string
	Copyright string
}

func (t *DefineFontName) Code() Code {
	return RecordDefineFontName
}
