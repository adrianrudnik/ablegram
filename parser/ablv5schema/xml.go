package ablv5schema

import "encoding/xml"

type Ableton struct {
	XMLName xml.Name `xml:"Ableton"`

	MajorVersion      int    `xml:"MajorVersion,attr"`
	MinorVersion      string `xml:"MinorVersion,attr"`
	SchemaChangeCount int    `xml:"SchemaChangeCount,attr"`
	Creator           string `xml:"Creator,attr"`
	Revision          string `xml:"Revision,attr"`

	LiveSet []LiveSet `xml:"LiveSet"`
}

type LiveSet struct {
	XMLName xml.Name `xml:"LiveSet"`
}
