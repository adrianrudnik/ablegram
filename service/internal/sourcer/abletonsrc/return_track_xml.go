package abletonsrc

type XmlReturnTrack struct {
	XmlIdNode
	XmlTrackNameNode
	XmlColorNode

	DeviceChain XmlTrackDeviceChain `xml:"DeviceChain"`
}
