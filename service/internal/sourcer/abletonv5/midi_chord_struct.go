package abletonv5

type MidiChordDeviceDocument struct {
	HasBase
	HasFileReference
	HasUserName
	HasUserInfoText
	HasDeviceIsExpanded
	HasDeviceIsFolded
}

func NewMidiChordDeviceDocument() *MidiChordDeviceDocument {
	return &MidiChordDeviceDocument{
		HasBase:             NewHasBase(AbletonMidiChordDevice),
		HasFileReference:    NewHasFileReference(),
		HasUserName:         NewHasUserName(),
		HasUserInfoText:     NewHasUserInfoText(),
		HasDeviceIsExpanded: NewHasDeviceIsExpanded(),
		HasDeviceIsFolded:   NewHasDeviceIsFolded(),
	}
}
