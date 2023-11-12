package abletonv5

type AudioTrackDocument struct {
	HasBase
	HasTrackUserNames
	HasIsFrozenOption
	HasColor
}

func NewAudioTrackDocument() *AudioTrackDocument {
	return &AudioTrackDocument{
		HasBase:           NewHasBase(AbletonAudioTrack),
		HasTrackUserNames: NewHasTrackUserNames(),
		HasIsFrozenOption: NewHasTrackIsFrozen(),
		HasColor:          NewHasColor(),
	}
}
