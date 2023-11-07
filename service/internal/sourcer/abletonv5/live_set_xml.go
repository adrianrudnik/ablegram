package abletonv5

import "encoding/xml"

type XmlLiveSet struct {
	XMLName          xml.Name                 `xml:"LiveSet"`
	Tracks           XmlTracks                `xml:"Tracks"`
	Scenes           []XmlScene               `xml:"Scenes"`
	ScaleInformation XmlScaleInformationValue `xml:"ScaleInformation"`
	InKey            XmlBooleanValue          `xml:"InKey"`
	MasterTrack      XmlMasterTrack           `xml:"MasterTrack"`
	Annotation       XmlStringValue           `xml:"Annotation"`
}

type XmlTracks struct {
	MidiTracks   []XmlMidiTrack    `xml:"MidiTrack"`
	AudioTracks  []XmlAudioTrack   `xml:"AudioTrack"`
	ReturnTracks []XmlReturnTrack  `xml:"ReturnTrack"`
	GroupTracks  []XmlGroupTrack   `xml:"GroupTrack"`
	PreHearTrack []XmlPreHearTrack `xml:"PreHearTrack"`
}

func (l *XmlLiveSet) GetAllTrackDeviceChains() []XmlTrackDeviceChain {
	hits := make([]XmlTrackDeviceChain, 0, 100)

	// Collect all possible device chains
	for _, midiTrack := range l.Tracks.MidiTracks {
		hits = append(hits, midiTrack.DeviceChain)
	}

	for _, audioTrack := range l.Tracks.AudioTracks {
		hits = append(hits, audioTrack.DeviceChain)
	}

	for _, returnTrack := range l.Tracks.ReturnTracks {
		hits = append(hits, returnTrack.DeviceChain)
	}

	for _, preHearTrack := range l.Tracks.PreHearTrack {
		hits = append(hits, preHearTrack.DeviceChain)
	}

	for _, groupTrack := range l.Tracks.GroupTracks {
		hits = append(hits, groupTrack.DeviceChain)
	}

	return hits
}

func (l *XmlLiveSet) GetAllActualDeviceChains() []XmlActualDeviceChain {
	hits := make([]XmlActualDeviceChain, 0, 100)

	hits = append(hits, l.MasterTrack.DeviceChain.DeviceChain)

	for _, track := range l.Tracks.MidiTracks {
		hits = append(hits, track.DeviceChain.DeviceChain)
	}

	for _, track := range l.Tracks.AudioTracks {
		hits = append(hits, track.DeviceChain.DeviceChain)
	}

	for _, track := range l.Tracks.ReturnTracks {
		hits = append(hits, track.DeviceChain.DeviceChain)
	}

	for _, track := range l.Tracks.PreHearTrack {
		hits = append(hits, track.DeviceChain.DeviceChain)
	}

	for _, track := range l.Tracks.GroupTracks {
		hits = append(hits, track.DeviceChain.DeviceChain)
	}

	return hits
}
