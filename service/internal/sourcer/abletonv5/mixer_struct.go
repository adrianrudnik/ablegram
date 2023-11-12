package abletonv5

type MixerDocument struct {
	HasBase
	HasUserName
	HasUserInfoText
	HasIsFoldedOption
	HasIsExpandedOption
}

func NewMixerDocument() *MixerDocument {
	return &MixerDocument{
		HasBase:             NewHasBase(AbletonMixer),
		HasUserName:         NewHasUserName(),
		HasUserInfoText:     NewHasUserInfoText(),
		HasIsFoldedOption:   NewHasDeviceIsFolded(),
		HasIsExpandedOption: NewHasDeviceIsExpanded(),
	}
}
