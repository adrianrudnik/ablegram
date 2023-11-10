package abletonv5

type AudioTrackDocument struct {
	HasBase
	HasFileReference
	HasTrackUserNames
	HasIsFrozenOption
	HasColor
}

func NewAudioTrackDocument() *AudioTrackDocument {
	return &AudioTrackDocument{
		HasBase:           NewHasBase(AbletonAudioTrack),
		HasFileReference:  NewHasFileReference(),
		HasTrackUserNames: NewHasTrackUserNames(),
		HasIsFrozenOption: NewHasTrackIsFrozen(),
		HasColor:          NewHasColor(),
	}
}
