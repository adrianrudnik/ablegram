package abletonv5

type MidiVelocityDeviceDocument struct {
	HasBase
	HasFileReference
	HasUserName
	HasUserInfoText
	HasIsExpandedOption
	HasIsFoldedOption
}

func NewMidiVelocityDeviceDocument() *MidiVelocityDeviceDocument {
	return &MidiVelocityDeviceDocument{
		HasBase:             NewHasBase(AbletonMidiVelocityDevice),
		HasFileReference:    NewHasFileReference(),
		HasUserName:         NewHasUserName(),
		HasUserInfoText:     NewHasUserInfoText(),
		HasIsExpandedOption: NewHasDeviceIsExpanded(),
		HasIsFoldedOption:   NewHasDeviceIsFolded(),
	}
}
