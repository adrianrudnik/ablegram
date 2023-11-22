package abletonsrc

type XmlAudioTrack struct {
	XmlIdNode
	XmlTrackNameNode
	XmlColorNode
	XmlIsFrozenNode

	DeviceChain XmlTrackDeviceChain `xml:"DeviceChain"`
}
