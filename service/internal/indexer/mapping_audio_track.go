package indexer

import (
	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/mapping"
)

type AudioTrackDocument struct {
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

func NewAudioTrackDocument() *AudioTrackDocument {
	return &AudioTrackDocument{
		T: "AudioTrack",
	}
}
func (d *AudioTrackDocument) Type() string {
	return d.T
}

// @see service/vendor/github.com/blevesearch/bleve/v2/mapping.go
func buildAudioTrackMapping() *mapping.DocumentMapping {
	m := bleve.NewDocumentMapping()
	m.AddFieldMappingsAt("type", mapping.NewKeywordFieldMapping())
	m.AddFieldMappingsAt("tags", mapping.NewKeywordFieldMapping())

	m.AddFieldMappingsAt("pathFolder", mapping.NewKeywordFieldMapping())
	m.AddFieldMappingsAt("pathAbsolute", mapping.NewKeywordFieldMapping())
	m.AddFieldMappingsAt("filename", mapping.NewKeywordFieldMapping())

	m.AddFieldMappingsAt("displayName", NewFulltextTextFieldMapping(true))
	m.AddFieldMappingsAt("effectiveName", NewFulltextTextFieldMapping(false))
	m.AddFieldMappingsAt("userName", NewFulltextTextFieldMapping(false))
	m.AddFieldMappingsAt("memorizedFirstClipName", NewFulltextTextFieldMapping(false))
	m.AddFieldMappingsAt("annotation", NewFulltextTextFieldMapping(true))

	m.AddFieldMappingsAt("color", mapping.NewNumericFieldMapping())

	return m
}
