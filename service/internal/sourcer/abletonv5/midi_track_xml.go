package abletonv5

type XmlMidiTrack struct {
	XmlIdNode
	XmlTrackNameNode
	XmlColorNode
	XmlIsFrozenNode

	DeviceChain XmlTrackDeviceChain `xml:"DeviceChain"`
}
