package abletonsrc

import (
	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/mapping"
)

func buildInfotextMapping() *mapping.DocumentMapping {
	m := bleve.NewDocumentMapping()

	createBaseMappings(m)
	createUserInfoTextMapping(m)

	m.AddFieldMappingsAt("parent", mapping.NewKeywordFieldMapping())

	return m
}
