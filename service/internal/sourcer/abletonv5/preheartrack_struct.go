package abletonv5

type PreHearTrackDocument struct {
	*HasBase
	*HasFileReference
	*HasTrackUserNames
	*HasColor
}

func NewPreHearTrackDocument() *PreHearTrackDocument {
	return &PreHearTrackDocument{
		HasBase:           NewHasBase(AbletonPreHearTrack),
		HasFileReference:  NewHasFileReference(),
		HasTrackUserNames: NewHasTrackUserNames(),
		HasColor:          NewHasColor(),
	}
}
