package abletonv5

type XmlScene struct {
	XmlIdNode
	XmlAnnotationNode
	XmlTempoWithToggleNode

	Name  XmlStringValue `xml:"Name"`
	Color XmlColorValue  `xml:"Color"`
}
