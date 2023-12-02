package tag

type DefineSceneAndFrameLabelData struct {
	_ struct{} `swfFlags:"root"`

	SceneCount uint32 `swfFlags:"encoded"`
	Scenes     []struct {
		Offset uint32 `swfFlags:"encoded"`
		Name   string
	} `swfCount:"SceneCount"`

	FrameLabelCount uint32 `swfFlags:"encoded"`
	FrameLabels     []struct {
		FrameNumber uint32 `swfFlags:"encoded"`
		Label       string
	} `swfCount:"FrameLabelCount"`
}

func (t *DefineSceneAndFrameLabelData) Code() Code {
	return RecordDefineSceneAndFrameLabelData
}
