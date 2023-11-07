package abletonv5

type XmlPreHearTrack struct {
	XmlIdNode
	XmlTrackNameNode
	XmlColorNode

	DeviceChain XmlTrackDeviceChain `xml:"DeviceChain"`
}
