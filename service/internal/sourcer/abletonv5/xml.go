package abletonv5

import (
	"encoding/xml"
)

type XmlRoot struct {
	XMLName xml.Name `xml:"Ableton"`

	MajorVersion      string `xml:"MajorVersion,attr"`
	MinorVersion      string `xml:"MinorVersion,attr"`
	SchemaChangeCount int64  `xml:"SchemaChangeCount,attr"`
	Creator           string `xml:"Creator,attr"`
	Revision          string `xml:"Revision,attr"`

	LiveSet XmlLiveSet `xml:"LiveSet"`
}

type XmlIdNode struct {
	Id int64 `xml:"Id,attr"`
}

type XmlTrackNameNode struct {
	Name XmlTrackNameValue `xml:"Name"`
}

type XmlTrackNameValue struct {
	XmlUserNameNode
	XmlAnnotationNode

	EffectiveName          XmlStringValue `xml:"EffectiveName"`
	MemorizedFirstClipName XmlStringValue `xml:"MemorizedFirstClipName"`
}

type XmlUserNameNode struct {
	UserName XmlStringValue `xml:"UserName"`
}

type XmlAnnotationNode struct {
	Annotation XmlStringValue `xml:"Annotation"`
}

type XmlTempoWithToggleNode struct {
	Tempo        XmlIntValue     `xml:"Tempo"`
	TempoEnabled XmlBooleanValue `xml:"TempoEnabled"`
}

type XmlTempo struct {
	Manual XmlFloatValue `xml:"Manual"`
}

type XmlStringValue struct {
	Value string `xml:"Value,attr"`
}

type XmlIntValue struct {
	Value float64 `xml:"Value,attr"`
}

type XmlFloatValue struct {
	Value float64 `xml:"Value,attr"`
}

type XmlBooleanValue struct {
	Value bool `xml:"Value,attr"`
}

type XmlColorValue struct {
	Value int16 `xml:"Value,attr"`
}

type XmlRootNoteValue struct {
	Value int64 `xml:"Value,attr"`
}

type XmlIsFoldedValue struct {
	IsFolded XmlBooleanValue `xml:"IsFolded"`
}

type XmlIsExpandedValue struct {
	IsExpanded XmlBooleanValue `xml:"IsExpanded"`
}

type XmlColorNode struct {
	Color XmlColorValue `xml:"Color"`
}

type XmlIsFrozenNode struct {
	IsFrozen XmlBooleanValue `xml:"Freeze"`
}
