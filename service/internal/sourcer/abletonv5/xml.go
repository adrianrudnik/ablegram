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

type XmlLiveSet struct {
	XMLName          xml.Name            `xml:"LiveSet"`
	Tracks           XmlTracks           `xml:"Tracks"`
	ScaleInformation XmlScaleInformation `xml:"ScaleInformation"`
	InKey            XmlBooleanValue     `xml:"InKey"`
	MasterTrack      XmlMasterTrack      `xml:"MasterTrack"`
	Annotation       XmlStringValue      `xml:"Annotation"`
}

type XmlTracks struct {
	MidiTracks   []XmlMidiTrack   `xml:"MidiTrack"`
	AudioTracks  []XmlAudioTrack  `xml:"AudioTrack"`
	ReturnTracks []XmlReturnTrack `xml:"ReturnTrack"`
	GroupTracks  []XmlGroupTrack  `xml:"GroupTrack"`
}

type XmlMidiTrack struct {
	Id     int64           `xml:"Id,attr"`
	Name   XmlFullName     `xml:"Name"`
	Color  XmlColorValue   `xml:"Color"`
	Frozen XmlBooleanValue `xml:"Freeze"`

	DeviceChain XmlDeviceChain `xml:"DeviceChain"`
}

type XmlAudioTrack struct {
	Id     int64           `xml:"Id,attr"`
	Name   XmlFullName     `xml:"Name"`
	Color  XmlColorValue   `xml:"Color"`
	Frozen XmlBooleanValue `xml:"Freeze"`

	DeviceChain XmlDeviceChain `xml:"DeviceChain"`
}

type XmlReturnTrack struct {
	Id    int64         `xml:"Id,attr"`
	Name  XmlFullName   `xml:"Name"`
	Color XmlColorValue `xml:"Color"`

	DeviceChain XmlDeviceChain `xml:"DeviceChain"`
}

type XmlGroupTrack struct {
	Id          int64          `xml:"Id,attr"`
	Name        XmlFullName    `xml:"Name"`
	DeviceChain XmlDeviceChain `xml:"DeviceChain"`
	Color       XmlColorValue  `xml:"Color"`
}

type XmlFullName struct {
	*XmlUserName
	*XmlAnnotation

	EffectiveName          XmlStringValue
	MemorizedFirstClipName XmlStringValue
}

type XmlUserName struct {
	UserName XmlStringValue
}

type XmlAnnotation struct {
	Annotation XmlStringValue
}

type XmlScaleInformation struct {
	RootNote XmlIntValue    `xml:"RootNote"`
	Name     XmlStringValue `xml:"Name"`
}

func (s *XmlScaleInformation) HumanizeRootNote() string {
	switch s.RootNote.Value {
	case 0:
		return "c"
	}

	return "unknown"
}

type XmlMasterTrack struct {
	DeviceChain XmlDeviceChain `xml:"DeviceChain"`
}

type XmlDeviceChain struct {
	Mixer   XmlMixer      `xml:"Mixer"`
	Devices XmlDeviceList `xml:"Devices"`
}

type XmlDeviceList struct {
	Reverb []XmlReverbDevice `xml:"Reverb"`
	Delay  []XmlDelayDevice  `xml:"Delay"`
}

func (dl *XmlDeviceList) GetCount() uint64 {
	return uint64(len(dl.Reverb) + len(dl.Delay))
}

type XmlReverbDevice struct {
	*XmlUserName
	*XmlAnnotation
}

type XmlDelayDevice struct {
	*XmlUserName
	*XmlAnnotation
}

type XmlMixer struct {
	*XmlUserName
	*XmlAnnotation
	Tempo XmlTempo `xml:"Tempo"`
}

type XmlTempo struct {
	Manual XmlFloatValue `xml:"Manual"`
}

type XmlStringValue struct {
	Value string `xml:"Value,attr"`
}

type XmlIntValue struct {
	Value int64 `xml:"Value,attr"`
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
