package abletonsrc

import (
	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/mapping"
)

func buildPreHearTrackMapping() *mapping.DocumentMapping {
	m := bleve.NewDocumentMapping()

	createSharedTrackMappings(m)

	return m
}
