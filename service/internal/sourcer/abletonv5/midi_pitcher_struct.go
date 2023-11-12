package abletonv5

type MidiPitcherDeviceDocument struct {
	HasBase
	HasUserName
	HasUserInfoText
	HasIsExpandedOption
	HasIsFoldedOption
}

func NewMidiPitcherDeviceDocument() *MidiPitcherDeviceDocument {
	return &MidiPitcherDeviceDocument{
		HasBase:             NewHasBase(AbletonMidiPitcherDevice),
		HasUserName:         NewHasUserName(),
		HasUserInfoText:     NewHasUserInfoText(),
		HasIsExpandedOption: NewHasDeviceIsExpanded(),
		HasIsFoldedOption:   NewHasDeviceIsFolded(),
	}
}
