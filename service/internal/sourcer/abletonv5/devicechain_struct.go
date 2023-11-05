package abletonv5

type DeviceChainDocument struct {
	*HasBase
	*HasFileReference

	DeviceCount uint64 `json:"device_count"`
}

func NewDeviceChainDocument() *DeviceChainDocument {
	return &DeviceChainDocument{
		HasBase:          NewHasBase(AbletonDeviceChain),
		HasFileReference: NewHasFileReference(),
	}
}
