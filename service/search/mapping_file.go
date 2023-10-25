package search

import (
	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/mapping"
)

func buildAlsFileMapping(options *SearchOptions) *mapping.DocumentMapping {
	filenameMapping := getBaseFilenameFieldMapping()

	m := bleve.NewDocumentMapping()
	m.AddFieldMappingsAt("filename", filenameMapping)

	return m
}
