package tag

type Code uint16

const (
	RecordEnd = Code(iota)
	RecordShowFrame
	RecordDefineShape
	_
	RecordPlaceObject
	RecordRemoveObject
	RecordDefineBits
	_
	RecordJPEGTables
	RecordSetBackgroundColor
	RecordDefineFont
	RecordDefineText
	RecordDoAction
	RecordDefineFontInfo
	RecordDefineSound
	RecordStartSound
	_
	_
	RecordSoundStreamHead
	RecordSoundStreamBlock
	RecordDefineBitsLossless
	RecordDefineBitsJPEG2
	RecordDefineShape2
	_
	RecordProtect
	_
	RecordPlaceObject2
	_
	RecordRemoveObject2
	_
	_
	_
	RecordDefineShape3
	RecordDefineText2
	_
	RecordDefineBitsJPEG3
	RecordDefineBitsLossless2
	_
	_
	RecordDefineSprite
	_
	_
	_
	RecordFrameLabel
	_
	RecordSoundStreamHead2
	RecordDefineMorphShape
	_
	RecordDefineFont2
	_
	_
	_
	_
	_
	_
	_
	RecordExportAssets
	RecordImportAssets
	_
	RecordDoInitAction
	_
	_
	RecordDefineFontInfo2
	_
	_
	_
	_
	_
	_
	RecordFileAttributes
	RecordPlaceObject3
	RecordImportAssets2
	_
	RecordDefineFontAlignZones
	_
	RecordDefineFont3
	RecordSymbolClass
	RecordMetadata
	RecordDefineScalingGrid
	_
	_
	_
	RecordDoABC
	RecordDefineShape4
	RecordDefineMorphShape2
	_
	RecordDefineSceneAndFrameLabelData
	_
	RecordDefineFontName
	RecordStartSound2
	RecordDefineBitsJPEG4
	RecordDefineFont4
	_
	RecordEnableTelemetry
	RecordPlaceObject4
)
