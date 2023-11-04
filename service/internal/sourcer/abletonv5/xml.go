package abletonv5

import (
	"encoding/xml"
)

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
	XMLName          xml.Name         `xml:"LiveSet"`
	Tracks           Tracks           `xml:"Tracks"`
	ScaleInformation ScaleInformation `xml:"ScaleInformation"`
	InKey            BooleanValue     `xml:"InKey"`
	MasterTrack      MasterTrack      `xml:"MasterTrack"`
	Annotation       StringValue      `xml:"Annotation"`
}

type Tracks struct {
	MidiTracks  []MidiTrack  `xml:"MidiTrack"`
	AudioTracks []AudioTrack `xml:"AudioTrack"`
}

type MidiTrack struct {
	Id     int64        `xml:"Id,attr"`
	Name   TrackNames   `xml:"Name"`
	Color  ColorValue   `xml:"Color"`
	Frozen BooleanValue `xml:"Freeze"`
}

type AudioTrack struct {
	Id     int64        `xml:"Id,attr"`
	Name   TrackNames   `xml:"Name"`
	Color  ColorValue   `xml:"Color"`
	Frozen BooleanValue `xml:"Freeze"`
}

type TrackNames struct {
	EffectiveName          StringValue
	UserName               StringValue
	Annotation             StringValue
	MemorizedFirstClipName StringValue
}

type ScaleInformation struct {
	RootNote IntValue    `xml:"RootNote"`
	Name     StringValue `xml:"Name"`
}

func (s *ScaleInformation) HumanizeRootNote() string {
	switch s.RootNote.Value {
	case 0:
		return "c"
	}

	return "unknown"
}

type MasterTrack struct {
	DeviceChain DeviceChain `xml:"DeviceChain"`
}

type DeviceChain struct {
	Mixer Mixer `xml:"Mixer"`
}

type Mixer struct {
	Tempo Tempo `xml:"Tempo"`
}

type Tempo struct {
	Manual FloatValue `xml:"Manual"`
}

type StringValue struct {
	Value string `xml:"Value,attr"`
}

type IntValue struct {
	Value int64 `xml:"Value,attr"`
}

type FloatValue struct {
	Value float64 `xml:"Value,attr"`
}

type BooleanValue struct {
	Value bool `xml:"Value,attr"`
}

type ColorValue struct {
	Value int16 `xml:"Value,attr"`
}

type RootNoteValue struct {
	Value int64 `xml:"Value,attr"`
}
