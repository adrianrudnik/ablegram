package abletonv5

type ReturnTrackDocument struct {
	*HasBase
	*HasFileReference
	*HasName
	*HasColor
}

func NewReturnTrackDocument() *ReturnTrackDocument {
	return &ReturnTrackDocument{
		HasBase:          &HasBase{T: AbletonReturnTrack},
		HasFileReference: &HasFileReference{},
		HasName:          &HasName{},
		HasColor:         &HasColor{},
	}
}
