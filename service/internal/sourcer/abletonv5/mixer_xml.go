package abletonv5

type XmlMixer struct {
	XmlUserNameNode
	XmlAnnotationNode
	Tempo XmlTempo `xml:"Tempo"`
}
