package abletonv5

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type XmlRoot struct {
	XMLName xml.Name `xml:"Ableton"`

	MajorVersion      string `xml:"MajorVersion,attr"`
	MinorVersion      string `xml:"MinorVersion,attr"`
	SchemaChangeCount int64  `xml:"SchemaChangeCount,attr"`
	Creator           string `xml:"Creator,attr"`
	Revision          string `xml:"Revision,attr"`

	Raw string `xml:",innerxml"`

	LiveSet XmlLiveSet `xml:"LiveSet"`
}

func (x *XmlRoot) IsMinorVersion(v int) bool {
	return strings.HasPrefix(x.MinorVersion, fmt.Sprintf("%d.", v))
}

type XmlRootNoteValue struct {
	Value int64 `xml:"Value,attr"`
}
