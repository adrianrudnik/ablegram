package abletonv5

type AudioTrackDocument struct {
	*HasBase
	*HasFileReference
	*HasTrackUserNames
	*HasColor

	Frozen bool `json:"frozen,omitempty"`
}

func NewAudioTrackDocument() *AudioTrackDocument {
	return &AudioTrackDocument{
		HasBase:           NewHasBase(AbletonAudioTrack),
		HasFileReference:  NewHasFileReference(),
		HasTrackUserNames: NewHasTrackUserNames(),
		HasColor:          NewHasColor(),
	}
}
