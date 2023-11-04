package abletonv5

type ClipDocument struct {
	T    string   `json:"type"`
	Tags []string `json:"tags,omitempty"`

	PathAbsolute string `json:"pathAbsolute,omitempty"`
	PathFolder   string `json:"pathFolder,omitempty"`
	Filename     string `json:"filename,omitempty"`
}

func NewClipDocument() *ClipDocument {
	return &ClipDocument{
		T: "Clip",
	}
}

func (d *ClipDocument) Type() string {
	return d.T
}
