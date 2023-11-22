package abletonsrc

type XmlMasterTrack struct {
	XmlTrackNameNode
	XmlColorNode

	DeviceChain XmlTrackDeviceChain `xml:"DeviceChain"`
}
