package abletonv5

type MidiVelocityDeviceDocument struct {
	HasBase
	HasFileReference
	HasUserName
	HasUserInfoText
	HasDeviceIsExpanded
	HasDeviceIsFolded
}

func NewMidiVelocityDeviceDocument() *MidiVelocityDeviceDocument {
	return &MidiVelocityDeviceDocument{
		HasBase:             NewHasBase(AbletonMidiVelocityDevice),
		HasFileReference:    NewHasFileReference(),
		HasUserName:         NewHasUserName(),
		HasUserInfoText:     NewHasUserInfoText(),
		HasDeviceIsExpanded: NewHasDeviceIsExpanded(),
		HasDeviceIsFolded:   NewHasDeviceIsFolded(),
	}
}
