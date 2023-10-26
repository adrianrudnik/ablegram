package ablv5schema

import "encoding/xml"

type Ableton struct {
	XMLName xml.Name `xml:"Ableton"`

	MajorVersion      string `xml:"MajorVersion,attr"`
	MinorVersion      string `xml:"MinorVersion,attr"`
	SchemaChangeCount int64  `xml:"SchemaChangeCount,attr"`
	Creator           string `xml:"Creator,attr"`
	Revision          string `xml:"Revision,attr"`

	LiveSet LiveSet `xml:"LiveSet"`
}

type LiveSet struct {
	XMLName xml.Name `xml:"LiveSet"`
	Tracks  Tracks   `xml:"Tracks"`
}

type Tracks struct {
	MidiTracks  []MidiTrack  `xml:"MidiTrack"`
	AudioTracks []AudioTrack `xml:"AudioTrack"`
}

type MidiTrack struct {
	Id    int64      `xml:"Id,attr"`
	Name  TrackNames `xml:"Name"`
	Color ColorValue `xml:"ColorValue"`
}

type AudioTrack struct {
	Id     int64        `xml:"Id,attr"`
	Name   TrackNames   `xml:"Name"`
	Color  ColorValue   `xml:"ColorValue"`
	Frozen BooleanValue `xml:"Freeze"`
}

type TrackNames struct {
	EffectiveName          StringValue
	UserName               StringValue
	Annotation             StringValue
	MemorizedFirstClipName StringValue
}

type StringValue struct {
	Value string `xml:"Value,attr"`
}

type IntValue struct {
	Value int64 `xml:"Value,attr"`
}

type BooleanValue struct {
	Value bool `xml:"Value,attr"`
}

type ColorValue struct {
	Value int64 `xml:"Value,attr"`
}
