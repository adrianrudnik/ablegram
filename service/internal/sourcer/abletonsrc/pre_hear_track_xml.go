package abletonsrc

type XmlPreHearTrack struct {
	XmlIdNode
	XmlTrackNameNode
	XmlColorNode

	DeviceChain XmlTrackDeviceChain `xml:"DeviceChain"`
}
