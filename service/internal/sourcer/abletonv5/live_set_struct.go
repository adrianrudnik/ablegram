package abletonv5

type LiveSetDocument struct {
	HasBase
	HasUserInfoText
	HasScaleInformation

	MajorVersion string `json:"majorVersion,omitempty"`
	MinorVersion string `json:"minorVersion,omitempty"`
	Creator      string `json:"creator,omitempty"`
	Revision     string `json:"revision,omitempty"`

	InKey bool  `json:"inKey,omitempty"`
	Tempo int64 `json:"bpm,omitempty"`

	MidiTrackCount  int `json:"midiTrackCount,omitempty"`
	AudioTrackCount int `json:"audioTrackCount,omitempty"`
}

func NewLiveSetDocument() *LiveSetDocument {
	return &LiveSetDocument{
		HasBase:             NewHasBase(AbletonLiveSet),
		HasUserInfoText:     NewHasUserInfoText(),
		HasScaleInformation: NewHasScaleInformation(),
	}
}
