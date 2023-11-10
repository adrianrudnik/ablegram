package abletonv5

type MidiChordDeviceDocument struct {
	HasBase
	HasFileReference
	HasUserName
	HasUserInfoText
	HasIsExpandedOption
	HasIsFoldedOption
}

func NewMidiChordDeviceDocument() *MidiChordDeviceDocument {
	return &MidiChordDeviceDocument{
		HasBase:             NewHasBase(AbletonMidiChordDevice),
		HasFileReference:    NewHasFileReference(),
		HasUserName:         NewHasUserName(),
		HasUserInfoText:     NewHasUserInfoText(),
		HasIsExpandedOption: NewHasDeviceIsExpanded(),
		HasIsFoldedOption:   NewHasDeviceIsFolded(),
	}
}
