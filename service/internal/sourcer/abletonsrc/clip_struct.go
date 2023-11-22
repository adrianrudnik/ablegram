package abletonsrc

type MidiClipDocument struct {
	HasBase
	HasUserName
	HasUserInfoText
	HasColor
	HasScaleInformation
	HasTimeSignature
}

func NewMidiClipDocument() *MidiClipDocument {
	return &MidiClipDocument{
		HasBase:             NewHasBase(AbletonMidiClip),
		HasUserName:         NewHasUserName(),
		HasUserInfoText:     NewHasUserInfoText(),
		HasColor:            NewHasColor(),
		HasScaleInformation: NewHasScaleInformation(),
		HasTimeSignature:    NewHasTimeSignature(),
	}
}

type AudioClipDocument struct {
	HasBase
	HasUserName
	HasUserInfoText
}

func NewAudioClipDocument() *AudioClipDocument {
	return &AudioClipDocument{
		HasBase:         NewHasBase(AbletonAudioClip),
		HasUserName:     NewHasUserName(),
		HasUserInfoText: NewHasUserInfoText(),
	}
}
