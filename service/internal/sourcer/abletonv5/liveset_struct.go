package abletonv5

type LiveSetDocument struct {
	T    string   `json:"type"`
	Tags []string `json:"tags,omitempty"`

	PathAbsolute string `json:"pathAbsolute,omitempty"`
	PathFolder   string `json:"pathFolder,omitempty"`
	Filename     string `json:"filename,omitempty"`

	DisplayName  string `json:"displayName,omitempty"`
	MajorVersion string `json:"majorVersion,omitempty"`
	MinorVersion string `json:"minorVersion,omitempty"`
	Creator      string `json:"creator,omitempty"`
	Revision     string `json:"revision,omitempty"`
	Annotation   string `json:"annotation,omitempty"`

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
		T: "LiveSet",
	}
}

func (d *LiveSetDocument) Type() string {
	return d.T
}
