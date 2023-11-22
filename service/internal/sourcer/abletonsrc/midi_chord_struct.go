package abletonsrc

type MidiChordDeviceDocument struct {
	HasBase
	HasUserName
	HasUserInfoText
	HasIsExpandedOption
	HasIsFoldedOption
}

func NewMidiChordDeviceDocument() *MidiChordDeviceDocument {
	return &MidiChordDeviceDocument{
		HasBase:             NewHasBase(AbletonMidiChordDevice),
		HasUserName:         NewHasUserName(),
		HasUserInfoText:     NewHasUserInfoText(),
		HasIsExpandedOption: NewHasDeviceIsExpanded(),
		HasIsFoldedOption:   NewHasDeviceIsFolded(),
	}
}
