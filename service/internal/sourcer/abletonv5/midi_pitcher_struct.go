package abletonv5

type MidiPitcherDeviceDocument struct {
	HasBase
	HasFileReference
	HasUserName
	HasUserInfoText
	HasDeviceIsExpanded
	HasDeviceIsFolded
}

func NewMidiPitcherDeviceDocument() *MidiPitcherDeviceDocument {
	return &MidiPitcherDeviceDocument{
		HasBase:             NewHasBase(AbletonMidiPitcherDevice),
		HasFileReference:    NewHasFileReference(),
		HasUserName:         NewHasUserName(),
		HasUserInfoText:     NewHasUserInfoText(),
		HasDeviceIsExpanded: NewHasDeviceIsExpanded(),
		HasDeviceIsFolded:   NewHasDeviceIsFolded(),
	}
}
