package abletonv5

// HasBase is the minimum struct to comply to for the indexer
type HasBase struct {
	T    string   `json:"type"`
	Tags []string `json:"tags,omitempty"`
}

func (b *HasBase) Type() string {
	return b.T
}

// HasFileReference represents a link to a file that contained the element
type HasFileReference struct {
	PathAbsolute string `json:"pathAbsolute,omitempty"`
	PathFolder   string `json:"pathFolder,omitempty"`
	Filename     string `json:"filename,omitempty"`
}

// HasName represents an element that can be named by the user.
type HasName struct {
	DisplayName            string `json:"displayName,omitempty"`
	EffectiveName          string `json:"effectiveName,omitempty"`
	UserName               string `json:"userName,omitempty"`
	MemorizedFirstClipName string `json:"memorizedFirstClipName,omitempty"`
	Annotation             string `json:"annotation,omitempty"`
}

// HasColor represents an element that can be colored by the user.
type HasColor struct {
	Color int16 `json:"color,omitempty"`
}
