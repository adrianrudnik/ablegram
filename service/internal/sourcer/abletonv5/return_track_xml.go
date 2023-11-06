package abletonv5

type XmlReturnTrack struct {
	Id    int64         `xml:"Id,attr"`
	Name  XmlFullName   `xml:"Name"`
	Color XmlColorValue `xml:"Color"`

	DeviceChain XmlTrackDeviceChain `xml:"DeviceChain"`
}
