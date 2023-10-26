package search

import (
	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/analysis/analyzer/keyword"
	"github.com/blevesearch/bleve/v2/mapping"
)

func getBaseTagFieldMapping() *mapping.FieldMapping {
	m := bleve.NewTextFieldMapping()
	m.Analyzer = keyword.Name

	return m
}

func getBaseFilenameFieldMapping() *mapping.FieldMapping {
	m := bleve.NewTextFieldMapping()
	m.Store = true

	return m
}

type NameVariantDocument struct {
	DisplayName            string `json:"display_name,omitempty"`
	EffectiveName          string `json:"effective_name,omitempty"`
	UserName               string `json:"user_name,omitempty"`
	Annotation             string `json:"annotation,omitempty"`
	MemorizedFirstClipName string `json:"memorizedFirstClipName,omitempty"`
}

func getBaseNameVariantMapping(options *SearchOptions) *mapping.DocumentMapping {
	name := bleve.NewDocumentMapping()

	displayName := bleve.NewTextFieldMapping()
	displayName.Store = true
	displayName.Index = false

	name.AddFieldMappingsAt("display_name", displayName)

	effectiveName := bleve.NewTextFieldMapping()
	effectiveName.Store = false

	name.AddFieldMappingsAt("effective_name", effectiveName)

	userName := bleve.NewTextFieldMapping()
	userName.Store = false

	name.AddFieldMappingsAt("user_name", userName)

	annotation := bleve.NewTextFieldMapping()
	annotation.Store = false

	name.AddFieldMappingsAt("annotation", annotation)

	memorizedFirstClipName := bleve.NewTextFieldMapping()
	memorizedFirstClipName.Store = false

	name.AddFieldMappingsAt("memorized_first_clip_name", memorizedFirstClipName)

	return name
}
