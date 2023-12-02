package tag

type EnableTelemetry struct {
	Reserved     uint16 `swfBits:",16"`
	PasswordHash [32]byte
}

func (t *EnableTelemetry) Code() Code {
	return RecordEnableTelemetry
}
