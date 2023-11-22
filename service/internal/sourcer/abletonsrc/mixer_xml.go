package abletonsrc

type XmlMixer struct {
	XmlUserNameNode
	XmlAnnotationNode
	XmlIsFoldedNode
	XmlIsExpandedNode

	Tempo XmlTempoNode `xml:"Tempo"`
}
