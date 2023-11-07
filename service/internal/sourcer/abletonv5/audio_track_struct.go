package abletonv5

type AudioTrackDocument struct {
	HasBase
	HasFileReference
	HasTrackUserNames
	HasTrackIsFrozen
	HasColor
}

func NewAudioTrackDocument() *AudioTrackDocument {
	return &AudioTrackDocument{
		HasBase:           NewHasBase(AbletonAudioTrack),
		HasFileReference:  NewHasFileReference(),
		HasTrackUserNames: NewHasTrackUserNames(),
		HasTrackIsFrozen:  NewHasTrackIsFrozen(),
		HasColor:          NewHasColor(),
	}
}
