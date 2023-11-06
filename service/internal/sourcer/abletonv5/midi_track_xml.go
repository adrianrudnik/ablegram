package abletonv5

type XmlMidiTrack struct {
	Id     int64           `xml:"Id,attr"`
	Name   XmlFullName     `xml:"Name"`
	Color  XmlColorValue   `xml:"Color"`
	Frozen XmlBooleanValue `xml:"Freeze"`

	DeviceChain XmlTrackDeviceChain `xml:"DeviceChain"`
}
