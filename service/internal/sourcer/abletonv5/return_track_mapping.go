package abletonv5

import (
	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/mapping"
)

func buildReturnTrackMapping() *mapping.DocumentMapping {
	m := bleve.NewDocumentMapping()

	createSharedTrackMappings(m)
	createFrozenMapping(m)

	return m
}
