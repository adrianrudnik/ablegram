package abletonv5

import (
	"encoding/xml"
	"fmt"
	"strconv"
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

func (x *XmlRoot) IsFromMinorVersion(v int) bool {
	mv, err := extractMajorVersionNumber(x.MinorVersion)
	if err != nil {
		return false
	}

	return mv >= v
}

func (x *XmlRoot) IsToMinorVersion(v int) bool {
	mv, err := extractMajorVersionNumber(x.MinorVersion)
	if err != nil {
		return false
	}

	return mv <= v
}

func extractMajorVersionNumber(v string) (int, error) {
	mv := v[:strings.IndexByte(v, '.')]

	i, err := strconv.Atoi(mv)
	if err != nil {
		Logger.Warn().Err(err).
			Str("version", v).
			Msg("Failed to parse minor XML version")

		return 0, err
	}

	return i, nil
}

type XmlRootNoteValue struct {
	Value int64 `xml:"Value,attr"`
}
