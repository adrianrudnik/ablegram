package abletonv5

type GroupTrackDocument struct {
	HasBase
	HasTrackUserNames
	HasColor
}

func NewGroupTrackDocument() *GroupTrackDocument {
	return &GroupTrackDocument{
		HasBase:           NewHasBase(AbletonGroupTrack),
		HasTrackUserNames: NewHasTrackUserNames(),
		HasColor:          NewHasColor(),
	}
}
