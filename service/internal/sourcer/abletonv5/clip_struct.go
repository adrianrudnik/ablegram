package abletonv5

type ClipDocument struct {
	HasBase
	HasFileReference
}

func NewClipDocument() *ClipDocument {
	return &ClipDocument{
		HasBase:          NewHasBase(AbletonClip),
		HasFileReference: NewHasFileReference(),
	}
}
