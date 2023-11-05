package abletonv5

type MidiTrackDocument struct {
	*HasBase
	*HasFileReference
	*HasName
	*HasColor

	Frozen bool `json:"frozen,omitempty"`
}

func NewMidiTrackDocument() *MidiTrackDocument {
	return &MidiTrackDocument{
		HasBase:          &HasBase{T: AbletonMidiTrack},
		HasFileReference: &HasFileReference{},
		HasName:          &HasName{},
		HasColor:         &HasColor{},
	}
}
