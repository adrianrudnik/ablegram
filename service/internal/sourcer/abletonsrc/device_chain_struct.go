package abletonsrc

type DeviceChainDocument struct {
	HasBase

	DeviceCount int `json:"device_count"`
}

func NewDeviceChainDocument() *DeviceChainDocument {
	return &DeviceChainDocument{
		HasBase: NewHasBase(AbletonDeviceChain),
	}
}
