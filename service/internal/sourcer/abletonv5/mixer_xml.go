package abletonv5

type XmlMixer struct {
	XmlUserNameNode
	XmlAnnotationNode
	Tempo XmlTempoNode `xml:"Tempo"`
}
