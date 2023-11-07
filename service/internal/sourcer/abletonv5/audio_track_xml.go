package abletonv5

type XmlAudioTrack struct {
	XmlIdNode
	XmlTrackNameNode
	XmlColorNode
	XmlIsFrozenNode

	DeviceChain XmlTrackDeviceChain `xml:"DeviceChain"`
}
