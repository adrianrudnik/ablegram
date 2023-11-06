package abletonv5

type MidiTrackDocument struct {
	HasBase
	HasFileReference
	HasTrackUserNames
	HasColor

	Frozen bool `json:"frozen,omitempty"`
}

func NewMidiTrackDocument() *MidiTrackDocument {
	return &MidiTrackDocument{
		HasBase:           NewHasBase(AbletonMidiTrack),
		HasFileReference:  NewHasFileReference(),
		HasTrackUserNames: NewHasTrackUserNames(),
		HasColor:          NewHasColor(),
	}
}
