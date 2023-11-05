package abletonv5

type GroupTrackDocument struct {
	*HasBase
	*HasFileReference
	*HasName
	*HasColor
}

func NewGroupTrackDocument() *GroupTrackDocument {
	return &GroupTrackDocument{
		HasBase:          &HasBase{T: AbletonGroupTrack},
		HasFileReference: &HasFileReference{},
		HasName:          &HasName{},
		HasColor:         &HasColor{},
	}
}
