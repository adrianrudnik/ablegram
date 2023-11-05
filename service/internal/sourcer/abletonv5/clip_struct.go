package abletonv5

type ClipDocument struct {
	*HasBase
	*HasFileReference
}

func NewClipDocument() *ClipDocument {
	return &ClipDocument{
		HasBase:          &HasBase{T: "AbletonClip"},
		HasFileReference: &HasFileReference{},
	}
}
