package abletonv5

type DeviceChainDocument struct {
	HasBase
	HasFileReference

	DeviceCount int `json:"device_count"`
}

func NewDeviceChainDocument() *DeviceChainDocument {
	return &DeviceChainDocument{
		HasBase:          NewHasBase(AbletonDeviceChain),
		HasFileReference: NewHasFileReference(),
	}
}
