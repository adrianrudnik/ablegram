package abletonv5

type XmlReturnTrack struct {
	XmlIdNode
	XmlTrackNameNode
	XmlColorNode

	DeviceChain XmlTrackDeviceChain `xml:"DeviceChain"`
}
