package abletonsrc

type XmlClipSlotList struct {
	ClipSlots []XmlClipSlotEnvelope `xml:"ClipSlot"`
}

type XmlClipSlotEnvelope struct {
	XmlIdNode

	MidiClip  *XmlMidiClip    `xml:"ClipSlot>Value>MidiClip"`
	AudioClip *XmlAudioClip   `xml:"ClipSlot>Value>AudioClip"`
	HasStop   XmlBooleanValue `xml:"HasStopButton"`
}

type XmlClipSlot struct {
	MidiClip XmlMidiClip `xml:"MidiClip"`
}

type XmlNotesNode struct {
	Notes *XmlNotes `xml:"Notes"`
}

type XmlNotes struct {
	KeyTracks []XmlKeyTrack `xml:"KeyTracks>KeyTrack"`
}

type XmlKeyTrack struct {
	//XmlIdNode

	X string `xml:",innerxml"`
	//Duration float64 `xml:"Duration,attr"`
	//Velocity int64   `xml:"Velocity,attr"`
	//
	//// Inner string `xml:",innerxml"`
	////Duration float64 `xml:"Notes>MidiNoteEvent,Duration,attr"`
	//////Velocity float64     `xml:"Notes>MidiNoteEvent>Velocity,attr"`
	MidiKey XmlMidiKey        `xml:"MidiKey"`
	Notes   []XmlKeyTrackNote `xml:"Notes>MidiNoteEvent"`
}

type XmlKeyTrackNote struct {
	Time              float64 `xml:"Time,attr"`
	Duration          float64 `xml:"Duration,attr"`
	Velocity          float64 `xml:"Velocity,attr"`
	VelocityDeviation float64 `xml:"VelocityDeviation,attr"`
	Probability       float64 `xml:"Probability,attr"`
	IsEnabled         bool    `xml:"IsEnabled,attr"`
}

type XmlMidiClip struct {
	XmlIdNode
	XmlNameNode
	XmlAnnotationNode
	XmlColorNode
	XmlScaleInformationNode
	XmlRemoteableTimeSignatureNode
	XmlNotesNode
}

type XmlAudioClip struct {
	XmlIdNode
	XmlNameNode
	XmlAnnotationNode
	XmlClipIsDisabledNode
	XmlClipIsWarpedNode
	// @todo why ColorIndex?
}
