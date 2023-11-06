package abletonv5

type XmlScene struct {
	Id XmlIntValue `xml:"Id"`
	XmlAnnotation
	XmlTempoWithToggle
	Name  XmlStringValue `xml:"Name"`
	Color XmlColorValue  `xml:"Color"`
}
