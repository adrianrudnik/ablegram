package abletonv5

type MidiClipDocument struct {
	HasBase
	HasFileReference
	HasUserName
	HasUserInfoText
	HasColor
	HasScaleInformation
}

func NewMidiClipDocument() *MidiClipDocument {
	return &MidiClipDocument{
		HasBase:             NewHasBase(AbletonMidiClip),
		HasFileReference:    NewHasFileReference(),
		HasUserName:         NewHasUserName(),
		HasUserInfoText:     NewHasUserInfoText(),
		HasColor:            NewHasColor(),
		HasScaleInformation: NewHasScaleInformation(),
	}
}

type AudioClipDocument struct {
	HasBase
	HasFileReference
	HasUserName
	HasUserInfoText
}

func NewAudioClipDocument() *AudioClipDocument {
	return &AudioClipDocument{
		HasBase:          NewHasBase(AbletonAudioClip),
		HasFileReference: NewHasFileReference(),
		HasUserName:      NewHasUserName(),
		HasUserInfoText:  NewHasUserInfoText(),
	}
}
