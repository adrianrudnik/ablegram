package abletonv5

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

type XmlMidiClip struct {
	XmlIdNode
	XmlNameNode
	XmlAnnotationNode
	XmlColorNode
	XmlScaleInformationNode
}

type XmlAudioClip struct {
	XmlIdNode
	XmlNameNode
	XmlAnnotationNode
	XmlClipIsDisabledNode
	XmlClipIsWarpedNode
	// @todo why ColorIndex?
}
