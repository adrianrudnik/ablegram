package abletonv5

type MidiArpeggiatorDeviceDocument struct {
	HasBase
	HasFileReference
	HasUserName
	HasUserInfoText
}

func NewMidiArpeggiatorDeviceDocument() *MidiArpeggiatorDeviceDocument {
	return &MidiArpeggiatorDeviceDocument{
		HasBase:          NewHasBase(AbletonMidiArpeggiatorDevice),
		HasFileReference: NewHasFileReference(),
		HasUserName:      NewHasUserName(),
		HasUserInfoText:  NewHasUserInfoText(),
	}
}
