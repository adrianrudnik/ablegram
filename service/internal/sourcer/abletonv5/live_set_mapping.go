package abletonv5

import (
	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/mapping"
)

func buildLiveSetMapping() *mapping.DocumentMapping {
	m := bleve.NewDocumentMapping()

	createBaseMappings(m)
	createUserInfoTextMapping(m)
	createScaleInformationMapping(m)

	m.AddFieldMappingsAt("majorVersion", mapping.NewKeywordFieldMapping())
	m.AddFieldMappingsAt("minorVersion", mapping.NewKeywordFieldMapping())
	m.AddFieldMappingsAt("creator", mapping.NewKeywordFieldMapping())

	m.AddFieldMappingsAt("inKey", mapping.NewBooleanFieldMapping())
	m.AddFieldMappingsAt("bpm", mapping.NewNumericFieldMapping())

	m.AddFieldMappingsAt("midiTrackCount", mapping.NewNumericFieldMapping())
	m.AddFieldMappingsAt("audioTrackCount", mapping.NewNumericFieldMapping())

	return m
}
