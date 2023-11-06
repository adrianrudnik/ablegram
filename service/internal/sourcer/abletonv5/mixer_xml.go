package abletonv5

type XmlMixer struct {
	XmlUserName
	XmlAnnotation
	Tempo XmlTempo `xml:"Tempo"`
}
