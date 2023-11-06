package abletonv5

type ReturnTrackDocument struct {
	HasBase
	HasFileReference
	HasTrackUserNames
	HasColor
}

func NewReturnTrackDocument() *ReturnTrackDocument {
	return &ReturnTrackDocument{
		HasBase:           NewHasBase(AbletonReturnTrack),
		HasFileReference:  NewHasFileReference(),
		HasTrackUserNames: NewHasTrackUserNames(),
		HasColor:          NewHasColor(),
	}
}
