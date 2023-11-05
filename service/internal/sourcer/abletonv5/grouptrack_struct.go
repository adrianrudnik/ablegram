package abletonv5

type GroupTrackDocument struct {
	*HasBase
	*HasFileReference
	*HasTrackUserNames
	*HasColor
}

func NewGroupTrackDocument() *GroupTrackDocument {
	return &GroupTrackDocument{
		HasBase:           NewHasBase(AbletonGroupTrack),
		HasFileReference:  NewHasFileReference(),
		HasTrackUserNames: NewHasTrackUserNames(),
		HasColor:          NewHasColor(),
	}
}
