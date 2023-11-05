package abletonv5

import (
	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/mapping"
)

func buildLiveSetMapping() *mapping.DocumentMapping {
	m := bleve.NewDocumentMapping()

	createSharedTrackMappings(m)

	return m
}
