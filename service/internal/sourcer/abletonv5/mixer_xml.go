package abletonv5

type XmlMixer struct {
	XmlUserNameNode
	XmlAnnotationNode
	XmlIsFoldedNode
	XmlIsExpandedNode

	Tempo XmlTempoNode `xml:"Tempo"`
}
