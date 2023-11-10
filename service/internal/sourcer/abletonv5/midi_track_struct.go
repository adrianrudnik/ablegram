package abletonv5

type MidiTrackDocument struct {
	HasBase
	HasFileReference
	HasTrackUserNames
	HasIsFrozenOption
	HasColor
}

func NewMidiTrackDocument() *MidiTrackDocument {
	return &MidiTrackDocument{
		HasBase:           NewHasBase(AbletonMidiTrack),
		HasFileReference:  NewHasFileReference(),
		HasTrackUserNames: NewHasTrackUserNames(),
		HasIsFrozenOption: NewHasTrackIsFrozen(),
		HasColor:          NewHasColor(),
	}
}
