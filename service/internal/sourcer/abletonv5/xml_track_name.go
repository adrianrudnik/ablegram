package abletonv5

type XmlTrackNameNode struct {
	Name XmlTrackNameValue `xml:"Name"`
}

type XmlTrackNameValue struct {
	XmlUserNameNode
	XmlAnnotationNode

	EffectiveName          XmlStringValue `xml:"EffectiveName"`
	MemorizedFirstClipName XmlStringValue `xml:"MemorizedFirstClipName"`
}
