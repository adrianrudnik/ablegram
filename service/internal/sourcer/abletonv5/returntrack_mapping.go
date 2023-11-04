package abletonv5

import (
	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/mapping"
)

func buildReturnTrackMapping() *mapping.DocumentMapping {
	m := bleve.NewDocumentMapping()
	m.AddFieldMappingsAt("type", mapping.NewKeywordFieldMapping())
	m.AddFieldMappingsAt("tags", mapping.NewKeywordFieldMapping())

	m.AddFieldMappingsAt("pathFolder", mapping.NewKeywordFieldMapping())
	m.AddFieldMappingsAt("pathAbsolute", mapping.NewKeywordFieldMapping())
	m.AddFieldMappingsAt("filename", mapping.NewKeywordFieldMapping())

	m.AddFieldMappingsAt("displayName", NewFulltextTextFieldMapping(true))
	m.AddFieldMappingsAt("effectiveName", NewFulltextTextFieldMapping(false))
	m.AddFieldMappingsAt("userName", NewFulltextTextFieldMapping(false))
	m.AddFieldMappingsAt("memorizedFirstClipName", NewFulltextTextFieldMapping(false))
	m.AddFieldMappingsAt("annotation", NewFulltextTextFieldMapping(true))

	m.AddFieldMappingsAt("color", mapping.NewNumericFieldMapping())

	return m
}
