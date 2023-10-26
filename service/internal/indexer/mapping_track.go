package indexer

import (
	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/mapping"
)

type AudioTrackDocument struct {
	Filename string `json:"filename"`
}

func buildAudioTrackMapping(options *SearchOptions) *mapping.DocumentMapping {
	nameVariantMapping := NewSearchableTextFieldMapping(options)

	m := bleve.NewDocumentMapping()
	m.AddFieldMappingsAt("name", nameVariantMapping)

	return m
}
