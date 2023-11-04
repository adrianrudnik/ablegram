package abletonv5

import (
	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/mapping"
)

func buildLiveSetMapping() *mapping.DocumentMapping {
	m := bleve.NewDocumentMapping()
	m.AddFieldMappingsAt("type", mapping.NewKeywordFieldMapping())
	m.AddFieldMappingsAt("tags", mapping.NewKeywordFieldMapping())

	m.AddFieldMappingsAt("pathFolder", mapping.NewKeywordFieldMapping())
	m.AddFieldMappingsAt("pathAbsolute", mapping.NewKeywordFieldMapping())
	m.AddFieldMappingsAt("filename", mapping.NewKeywordFieldMapping())

	m.AddFieldMappingsAt("displayName", NewFulltextTextFieldMapping(true))
	m.AddFieldMappingsAt("filename", NewFileFieldMapping())
	m.AddFieldMappingsAt("path", NewFulltextTextFieldMapping(false))
	m.AddFieldMappingsAt("majorVersion", NewPayloadFieldMapping())
	m.AddFieldMappingsAt("minorVersion", NewPayloadFieldMapping())
	m.AddFieldMappingsAt("creator", mapping.NewKeywordFieldMapping())
	m.AddFieldMappingsAt("revision", NewPayloadFieldMapping())
	m.AddFieldMappingsAt("annotation", NewFulltextTextFieldMapping(true))

	m.AddFieldMappingsAt("scale", mapping.NewKeywordFieldMapping())
	m.AddFieldMappingsAt("scaleName", mapping.NewKeywordFieldMapping())
	m.AddFieldMappingsAt("scaleRootNote", mapping.NewKeywordFieldMapping())

	m.AddFieldMappingsAt("inKey", mapping.NewBooleanFieldMapping())
	m.AddFieldMappingsAt("bpm", mapping.NewNumericFieldMapping())

	m.AddFieldMappingsAt("midiTrackCount", mapping.NewNumericFieldMapping())
	m.AddFieldMappingsAt("audioTrackCount", mapping.NewNumericFieldMapping())

	return m
}
