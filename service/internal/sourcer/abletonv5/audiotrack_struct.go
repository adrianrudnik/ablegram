package abletonv5

type AudioTrackDocument struct {
	*HasBase
	*HasFileReference
	*HasName
	*HasColor

	Frozen bool `json:"frozen,omitempty"`
}

func NewAudioTrackDocument() *AudioTrackDocument {
	return &AudioTrackDocument{
		HasBase:          &HasBase{T: AbletonAudioTrack},
		HasFileReference: &HasFileReference{},
		HasName:          &HasName{},
		HasColor:         &HasColor{},
	}
}
