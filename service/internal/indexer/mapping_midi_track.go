package indexer

import (
	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/mapping"
)

type MidiTrackDocument struct {
	T    string   `json:"type"`
	Tags []string `json:"tags,omitempty"`

	DisplayName            string `json:"display_name,omitempty"`
	EffectiveName          string `json:"effective_name,omitempty"`
	UserName               string `json:"user_name,omitempty"`
	Annotation             string `json:"annotation,omitempty"`
	MemorizedFirstClipName string `json:"memorized_first_clip_name,omitempty"`
	Filename               string `json:"filename,omitempty"`
}

func NewMidiTrackDocument() *MidiTrackDocument {
	return &MidiTrackDocument{
		T: "MidiTrack",
	}
}
func (d *MidiTrackDocument) Type() string {
	return d.T
}

// @see service/vendor/github.com/blevesearch/bleve/v2/mapping.go
func buildMidiTrackMapping(options *SearchOptions) *mapping.DocumentMapping {
	displayNameMapping := NewSearchableTextFieldMapping(options)
	displayNameMapping.Index = false

	hiddenNameMapping := NewSearchableTextFieldMapping(options)
	hiddenNameMapping.Store = false

	m := bleve.NewDocumentMapping()
	m.AddFieldMappingsAt("type", NewTypeFieldMapping(options))
	m.AddFieldMappingsAt("tags", NewTagFieldMapping(options))
	m.AddFieldMappingsAt("display_name", displayNameMapping)
	m.AddFieldMappingsAt("effective_name", hiddenNameMapping)
	m.AddFieldMappingsAt("user_name", hiddenNameMapping)
	m.AddFieldMappingsAt("annotation", hiddenNameMapping)
	m.AddFieldMappingsAt("memorized_first_clip_name", hiddenNameMapping)
	m.AddFieldMappingsAt("filename", NewFileFieldMapping(options))

	return m
}
