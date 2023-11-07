package abletonv5

type MixerDocument struct {
	HasBase
	HasFileReference
	HasUserName
	HasUserInfoText
	HasDeviceIsFolded
	HasDeviceIsExpanded
}

func NewMixerDocument() *MixerDocument {
	return &MixerDocument{
		HasBase:             NewHasBase(AbletonMixer),
		HasFileReference:    NewHasFileReference(),
		HasUserName:         NewHasUserName(),
		HasUserInfoText:     NewHasUserInfoText(),
		HasDeviceIsFolded:   NewHasDeviceIsFolded(),
		HasDeviceIsExpanded: NewHasDeviceIsExpanded(),
	}
}
