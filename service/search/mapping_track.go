package search

import (
	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/mapping"
)

type AudioTrackDocument struct {
	Name     NameVariantDocument `json:"name_variants"`
	Filename string              `json:"filename"`
}

func buildAudioTrackMapping(options *SearchOptions) *mapping.DocumentMapping {
	nameVariantMapping := getBaseNameVariantMapping(options)

	m := bleve.NewDocumentMapping()
	m.AddSubDocumentMapping("name_variants", nameVariantMapping)

	return m
}

type MidiTrackDocument struct {
	Name     NameVariantDocument `json:"name_variants"`
	Filename string              `json:"filename"`
}

func buildMidiTrackMapping(options *SearchOptions) *mapping.DocumentMapping {
	nameVariantMapping := getBaseNameVariantMapping(options)

	m := bleve.NewDocumentMapping()
	m.AddSubDocumentMapping("name_variants", nameVariantMapping)

	return m
}
