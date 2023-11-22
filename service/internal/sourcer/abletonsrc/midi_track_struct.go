package abletonsrc

type MidiTrackDocument struct {
	HasBase
	HasTrackUserNames
	HasIsFrozenOption
	HasColor
}

func NewMidiTrackDocument() *MidiTrackDocument {
	return &MidiTrackDocument{
		HasBase:           NewHasBase(AbletonMidiTrack),
		HasTrackUserNames: NewHasTrackUserNames(),
		HasIsFrozenOption: NewHasTrackIsFrozen(),
		HasColor:          NewHasColor(),
	}
}
