package abletonv5

type XmlPreHearTrack struct {
	Id          int64               `xml:"Id,attr"`
	Name        XmlFullName         `xml:"Name"`
	Color       XmlColorValue       `xml:"Color"`
	DeviceChain XmlTrackDeviceChain `xml:"DeviceChain"`
}
