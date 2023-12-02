package tag

import (
	"git.gammaspectra.live/WeebDataHoarder/swf-go/types"
)

type SoundFormat uint8

const (
	SoundFormatPCMNative = SoundFormat(iota)
	SoundFormatADPCM
	SoundFormatMP3
	SoundFormatPCMLittle
	SoundFormatNellymoser16KHz
	SoundFormatNellymoser8KHz
	SoundFormatNellymoser
	_
	_
	_
	_
	SoundFormatSpeex
)

type SoundRate uint8

const (
	SoundRate5512Hz = SoundRate(iota)
	SoundRate11025Hz
	SoundRate22050Hz
	SoundRate44100Hz
)

type SOUNDINFO struct {
	_    struct{} `swfFlags:"root"`
	Flag struct {
		Reserved       uint8 `swfBits:",2"`
		SyncStop       bool
		SyncNoMultiple bool
		HasEnvelope    bool
		HasLoops       bool
		HasOutPoint    bool
		HasInPoint     bool
	}
	InPoint         uint32          `swfCondition:"Flag.HasInPoint"`
	OutPoint        uint32          `swfCondition:"Flag.HasOutPoint"`
	LoopCount       uint16          `swfCondition:"Flag.HasLoops"`
	EnvPoints       uint8           `swfCondition:"Flag.HasEnvelope"`
	EnvelopeRecords []SOUNDENVELOPE `swfCondition:"Flag.HasEnvelope" swfCount:"EnvPoints"`
}

type SOUNDENVELOPE struct {
	Pos44      uint32
	LeftLevel  uint16
	RightLevel uint16
}

type StartSound struct {
	_         struct{} `swfFlags:"root"`
	SoundId   uint16
	SoundInfo SOUNDINFO
}

func (t *StartSound) Code() Code {
	return RecordStartSound
}

type StartSound2 struct {
	_              struct{} `swfFlags:"root"`
	SoundId        uint16
	SoundClassName string
	SoundInfo      SOUNDINFO
}

func (t *StartSound2) Code() Code {
	return RecordStartSound2
}

type DefineSound struct {
	_                struct{} `swfFlags:"root"`
	SoundId          uint16
	SoundFormat      SoundFormat `swfBits:",4"`
	SoundRate        SoundRate   `swfBits:",2"`
	SoundSize        uint8       `swfBits:",1"`
	IsStereo         bool
	SoundSampleCount uint32
	SoundData        types.UntilEndBytes
}

func (t *DefineSound) Code() Code {
	return RecordDefineSound
}

type SoundCompression uint8

const (
	_ = SoundCompression(iota)
	SoundCompressionADPCM
	SoundCompressionMP3
)

type SoundStreamHead struct {
	_                      struct{}  `swfFlags:"root"`
	Reserved               uint8     `swfBits:",4"`
	PlaybackSoundRate      SoundRate `swfBits:",2"`
	PlaybackSoundSize      uint8     `swfBits:",1"`
	PlaybackIsStereo       bool
	StreamSoundCompression SoundCompression `swfBits:",4"`
	StreamSoundRate        SoundRate        `swfBits:",2"`
	StreamSoundSize        uint8            `swfBits:",1"`
	StreamIsStereo         bool
	StreamSampleCount      uint16
	LatencySeek            int16 `swfCondition:"HasLatencySeek()"`
}

func (t *SoundStreamHead) HasLatencySeek(ctx types.ReaderContext) bool {
	return t.StreamSoundCompression == SoundCompressionMP3
}

func (t *SoundStreamHead) Code() Code {
	return RecordSoundStreamHead
}

type SoundStreamHead2 struct {
	_                 struct{}  `swfFlags:"root"`
	Reserved          uint8     `swfBits:",4"`
	PlaybackSoundRate SoundRate `swfBits:",2"`
	PlaybackSoundSize uint8     `swfBits:",1"`
	PlaybackIsStereo  bool
	StreamSoundFormat SoundFormat `swfBits:",4"`
	StreamSoundRate   SoundRate   `swfBits:",2"`
	StreamSoundSize   uint8       `swfBits:",1"`
	StreamIsStereo    bool
	StreamSampleCount uint16
	LatencySeek       int16 `swfCondition:"HasLatencySeek()"`
}

func (t *SoundStreamHead2) HasLatencySeek(ctx types.ReaderContext) bool {
	return t.StreamSoundFormat == SoundFormatMP3
}

func (t *SoundStreamHead2) Code() Code {
	return RecordSoundStreamHead2
}

type SoundStreamBlock struct {
	_    struct{} `swfFlags:"root"`
	Data types.UntilEndBytes
}

func (t *SoundStreamBlock) Code() Code {
	return RecordSoundStreamBlock
}
