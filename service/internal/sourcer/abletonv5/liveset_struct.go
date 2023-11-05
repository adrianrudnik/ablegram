package abletonv5

type LiveSetDocument struct {
	*HasBase
	*HasFileReference
	*HasUserInfoText

	DisplayName  string `json:"displayName,omitempty"`
	MajorVersion string `json:"majorVersion,omitempty"`
	MinorVersion string `json:"minorVersion,omitempty"`
	Creator      string `json:"creator,omitempty"`
	Revision     string `json:"revision,omitempty"`

	Scale         string `json:"scale,omitempty"`
	ScaleName     string `json:"scaleName,omitempty"`
	ScaleRootNote string `json:"scaleRootNote,omitempty"`

	InKey bool  `json:"inKey,omitempty"`
	Tempo int64 `json:"bpm,omitempty"`

	MidiTrackCount  int `json:"midiTrackCount,omitempty"`
	AudioTrackCount int `json:"audioTrackCount,omitempty"`

	ScaleInformation string ``
}

func NewLiveSetDocument() *LiveSetDocument {
	return &LiveSetDocument{
		HasBase:          NewHasBase(AbletonLiveSet),
		HasFileReference: NewHasFileReference(),
	}
}
