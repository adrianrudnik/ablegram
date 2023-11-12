package abletonv5

type ReturnTrackDocument struct {
	HasBase
	HasTrackUserNames
	HasColor
}

func NewReturnTrackDocument() *ReturnTrackDocument {
	return &ReturnTrackDocument{
		HasBase:           NewHasBase(AbletonReturnTrack),
		HasTrackUserNames: NewHasTrackUserNames(),
		HasColor:          NewHasColor(),
	}
}
