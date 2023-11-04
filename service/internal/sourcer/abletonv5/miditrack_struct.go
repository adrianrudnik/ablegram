package abletonv5

type MidiTrackDocument struct {
	T    string   `json:"type"`
	Tags []string `json:"tags,omitempty"`

	PathAbsolute string `json:"pathAbsolute,omitempty"`
	PathFolder   string `json:"pathFolder,omitempty"`
	Filename     string `json:"filename,omitempty"`

	DisplayName            string `json:"displayName,omitempty"`
	EffectiveName          string `json:"effectiveName,omitempty"`
	UserName               string `json:"userName,omitempty"`
	MemorizedFirstClipName string `json:"memorizedFirstClipName,omitempty"`
	Annotation             string `json:"annotation,omitempty"`

	Color  int16 `json:"color,omitempty"`
	Frozen bool  `json:"frozen,omitempty"`
}

func NewMidiTrackDocument() *MidiTrackDocument {
	return &MidiTrackDocument{
		T: "MidiTrack",
	}
}

func (d *MidiTrackDocument) Type() string {
	return d.T
}
