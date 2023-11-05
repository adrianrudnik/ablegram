package abletonv5

import (
	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/mapping"
)

func NewTypeFieldMapping() *mapping.FieldMapping {
	return bleve.NewKeywordFieldMapping()
}

func NewTagFieldMapping() *mapping.FieldMapping {
	return bleve.NewKeywordFieldMapping()
}

func NewFileFieldMapping() *mapping.FieldMapping {
	return bleve.NewKeywordFieldMapping()
}

func NewFulltextTextFieldMapping(store bool) *mapping.FieldMapping {
	fm := bleve.NewTextFieldMapping()
	fm.Analyzer = "enWithEdgeNgram325"
	fm.Store = store
	return fm
}

func NewPayloadFieldMapping() *mapping.FieldMapping {
	fm := bleve.NewTextFieldMapping()
	fm.Store = true
	fm.Index = false
	return fm
}

func NewInfoTextFieldMapping() *mapping.FieldMapping {
	fm := bleve.NewKeywordFieldMapping()
	fm.Store = true
	fm.Index = false
	return fm
}

func createBaseMappings(im *mapping.DocumentMapping) {
	im.AddFieldMappingsAt("type", mapping.NewKeywordFieldMapping())
	im.AddFieldMappingsAt("tags", mapping.NewKeywordFieldMapping())
}

func createFileReferenceMappings(im *mapping.DocumentMapping) {
	im.AddFieldMappingsAt("pathFolder", mapping.NewKeywordFieldMapping())
	im.AddFieldMappingsAt("pathAbsolute", mapping.NewKeywordFieldMapping())
	im.AddFieldMappingsAt("filename", mapping.NewKeywordFieldMapping())
}

func createUserNameMapping(im *mapping.DocumentMapping) {
	im.AddFieldMappingsAt("userName", mapping.NewKeywordFieldMapping())
}

func createUserInfoTextMapping(im *mapping.DocumentMapping) {
	im.AddFieldMappingsAt("annotation", mapping.NewTextFieldMapping())
}

func createFullNameMapping(im *mapping.DocumentMapping) {
	createUserNameMapping(im)
	createUserInfoTextMapping(im)

	im.AddFieldMappingsAt("displayName", NewFulltextTextFieldMapping(true))
	im.AddFieldMappingsAt("effectiveName", NewFulltextTextFieldMapping(false))
	im.AddFieldMappingsAt("memorizedFirstClipName", NewFulltextTextFieldMapping(false))
}

func createSharedTrackMappings(im *mapping.DocumentMapping) {
	createBaseMappings(im)
	createFileReferenceMappings(im)
	createFullNameMapping(im)
	createColorMapping(im)
}

func createFrozenMapping(im *mapping.DocumentMapping) {
	im.AddFieldMappingsAt("frozen", mapping.NewBooleanFieldMapping())
}

func createColorMapping(im *mapping.DocumentMapping) {
	im.AddFieldMappingsAt("color", mapping.NewNumericFieldMapping())
}

func createTempoWithToggle(im *mapping.DocumentMapping) {
	im.AddFieldMappingsAt("tempo", mapping.NewNumericFieldMapping())
	im.AddFieldMappingsAt("tempoEnabled", mapping.NewBooleanFieldMapping())
}
