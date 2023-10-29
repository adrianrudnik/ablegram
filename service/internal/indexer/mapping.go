package indexer

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
