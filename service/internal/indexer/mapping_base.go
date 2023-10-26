package indexer

import (
	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/mapping"
)

func NewTypeFieldMapping(options *SearchOptions) *mapping.FieldMapping {
	return bleve.NewKeywordFieldMapping()
}

func NewTagFieldMapping(options *SearchOptions) *mapping.FieldMapping {
	return bleve.NewKeywordFieldMapping()
}

func NewFileFieldMapping(options *SearchOptions) *mapping.FieldMapping {
	return bleve.NewKeywordFieldMapping()
}

func NewSearchableTextFieldMapping(options *SearchOptions) *mapping.FieldMapping {
	fm := bleve.NewTextFieldMapping()
	fm.Analyzer = "enWithEdgeNgram325"
	return fm
}

func NewInfoTextFieldMapping(options *SearchOptions) *mapping.FieldMapping {
	fm := bleve.NewKeywordFieldMapping()
	fm.Store = true
	fm.Index = false
	return fm
}
