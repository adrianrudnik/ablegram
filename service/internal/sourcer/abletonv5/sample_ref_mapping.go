package abletonv5

import (
	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/mapping"
)

func buildSampleReferenceMapping() *mapping.DocumentMapping {
	m := bleve.NewDocumentMapping()

	createBaseMappings(m)
	createFileReferenceMappings(m)

	m.AddFieldMappingsAt("sampleAbsPath", mapping.NewKeywordFieldMapping())
	m.AddFieldMappingsAt("sampleFilename", mapping.NewKeywordFieldMapping())
	m.AddFieldMappingsAt("sampleOriginalFileSize", mapping.NewNumericFieldMapping())

	return m
}
