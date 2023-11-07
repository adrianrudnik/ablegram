package abletonv5

type MidiTrackDocument struct {
	HasBase
	HasFileReference
	HasTrackUserNames
	HasTrackIsFrozen
	HasColor
}

func NewMidiTrackDocument() *MidiTrackDocument {
	return &MidiTrackDocument{
		HasBase:           NewHasBase(AbletonMidiTrack),
		HasFileReference:  NewHasFileReference(),
		HasTrackUserNames: NewHasTrackUserNames(),
		HasTrackIsFrozen:  NewHasTrackIsFrozen(),
		HasColor:          NewHasColor(),
	}
}
