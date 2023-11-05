package abletonv5

type GroupTrackDocument struct {
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

	Color int16 `json:"color,omitempty"`
}

func NewGroupTrackDocument() *GroupTrackDocument {
	return &GroupTrackDocument{
		T: "GroupTrack",
	}
}

func (d *GroupTrackDocument) Type() string {
	return d.T
}
