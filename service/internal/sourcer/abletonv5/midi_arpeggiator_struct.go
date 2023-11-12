package abletonv5

type MidiArpeggiatorDeviceDocument struct {
	HasBase
	HasUserName
	HasUserInfoText
}

func NewMidiArpeggiatorDeviceDocument() *MidiArpeggiatorDeviceDocument {
	return &MidiArpeggiatorDeviceDocument{
		HasBase:         NewHasBase(AbletonMidiArpeggiatorDevice),
		HasUserName:     NewHasUserName(),
		HasUserInfoText: NewHasUserInfoText(),
	}
}
