package abletonsrc

type PreHearTrackDocument struct {
	HasBase
	HasTrackUserNames
	HasColor
}

func NewPreHearTrackDocument() *PreHearTrackDocument {
	return &PreHearTrackDocument{
		HasBase:           NewHasBase(AbletonPreHearTrack),
		HasTrackUserNames: NewHasTrackUserNames(),
		HasColor:          NewHasColor(),
	}
}
