package tag

type FileAttributes struct {
	Reserved1     uint8 `swfBits:",1"`
	UseDirectBlit bool
	UseGPU        bool
	HasMetadata   bool
	ActionScript3 bool
	Reserved2     uint8 `swfBits:",2"`
	UseNetwork    bool
	Reserved      uint32 `swfBits:",24"`
}

func (s *FileAttributes) Code() Code {
	return RecordFileAttributes
}
