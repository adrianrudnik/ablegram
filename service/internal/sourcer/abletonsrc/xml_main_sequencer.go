package abletonsrc

type XmlMainSequencerNode struct {
	XmlUserNameNode
	XmlAnnotationNode
	XmlIsExpandedNode
	XmlIsFoldedNode
	ClipSlotList XmlClipSlotList `xml:"ClipSlotList"`
}
