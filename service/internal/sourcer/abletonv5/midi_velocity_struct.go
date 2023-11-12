package abletonv5

type MidiVelocityDeviceDocument struct {
	HasBase
	HasUserName
	HasUserInfoText
	HasIsExpandedOption
	HasIsFoldedOption
}

func NewMidiVelocityDeviceDocument() *MidiVelocityDeviceDocument {
	return &MidiVelocityDeviceDocument{
		HasBase:             NewHasBase(AbletonMidiVelocityDevice),
		HasUserName:         NewHasUserName(),
		HasUserInfoText:     NewHasUserInfoText(),
		HasIsExpandedOption: NewHasDeviceIsExpanded(),
		HasIsFoldedOption:   NewHasDeviceIsFolded(),
	}
}
