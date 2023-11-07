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

type XmlRootNoteValue struct {
	Value int64 `xml:"Value,attr"`
}
