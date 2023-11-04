package abletonv5

import (
	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/mapping"
)

func buildClipMapping() *mapping.DocumentMapping {
	m := bleve.NewDocumentMapping()
	m.AddFieldMappingsAt("type", mapping.NewKeywordFieldMapping())
	m.AddFieldMappingsAt("tags", mapping.NewKeywordFieldMapping())

	m.AddFieldMappingsAt("pathFolder", mapping.NewKeywordFieldMapping())
	m.AddFieldMappingsAt("pathAbsolute", mapping.NewKeywordFieldMapping())
	m.AddFieldMappingsAt("filename", mapping.NewKeywordFieldMapping())

	return m
}
