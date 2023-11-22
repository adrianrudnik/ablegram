package abletonsrc

type XmlGroupTrack struct {
	XmlIdNode
	XmlTrackNameNode

	Color       XmlColorValue       `xml:"Color"`
	DeviceChain XmlTrackDeviceChain `xml:"DeviceChain"`
}
