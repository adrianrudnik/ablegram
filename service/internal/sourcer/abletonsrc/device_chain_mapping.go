package abletonsrc

import (
	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/mapping"
)

func buildDeviceChainMapping() *mapping.DocumentMapping {
	m := bleve.NewDocumentMapping()

	createBaseMappings(m)

	m.AddFieldMappingsAt("device_count", mapping.NewNumericFieldMapping())

	return m
}
