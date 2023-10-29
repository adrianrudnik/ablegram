package indexer

import (
	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/mapping"
)

type AudioTrackDocument struct {
	T    string   `json:"type"`
	Tags []string `json:"tags,omitempty"`

	DisplayName            string `json:"display_name,omitempty"`
	EffectiveName          string `json:"effective_name,omitempty"`
	UserName               string `json:"user_name,omitempty"`
	Annotation             string `json:"annotation,omitempty"`
	MemorizedFirstClipName string `json:"memorized_first_clip_name,omitempty"`
	Filename               string `json:"filename,omitempty"`
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

	m.AddFieldMappingsAt("displayName", NewPayloadFieldMapping())
	m.AddFieldMappingsAt("effectiveName", NewFulltextTextFieldMapping(false))
	m.AddFieldMappingsAt("userName", NewFulltextTextFieldMapping(false))
	m.AddFieldMappingsAt("annotation", NewFulltextTextFieldMapping(false))
	m.AddFieldMappingsAt("memorizedFirstClipName", NewFulltextTextFieldMapping(false))
	m.AddFieldMappingsAt("filename", mapping.NewKeywordFieldMapping())

	return m
}
