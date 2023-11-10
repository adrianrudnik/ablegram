package abletonv5

type MidiPitcherDeviceDocument struct {
	HasBase
	HasFileReference
	HasUserName
	HasUserInfoText
	HasIsExpandedOption
	HasIsFoldedOption
}

func NewMidiPitcherDeviceDocument() *MidiPitcherDeviceDocument {
	return &MidiPitcherDeviceDocument{
		HasBase:             NewHasBase(AbletonMidiPitcherDevice),
		HasFileReference:    NewHasFileReference(),
		HasUserName:         NewHasUserName(),
		HasUserInfoText:     NewHasUserInfoText(),
		HasIsExpandedOption: NewHasDeviceIsExpanded(),
		HasIsFoldedOption:   NewHasDeviceIsFolded(),
	}
}
