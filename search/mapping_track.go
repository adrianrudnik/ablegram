package search

import (
	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/mapping"
)

func createAudioTrackMapping(options *SearchOptions) *mapping.DocumentMapping {
	trackMapping := bleve.NewDocumentMapping()

	trackMapping.AddSubDocumentMapping("name_variants", createNameMapping(options))

	return trackMapping
}

func createNameMapping(options *SearchOptions) *mapping.DocumentMapping {
	name := bleve.NewDocumentMapping()

	effectiveName := bleve.NewTextFieldMapping()
	effectiveName.Store = false
	effectiveName.Analyzer = options.PrimaryLanguage

	name.AddFieldMappingsAt("effective_name", effectiveName)

	userName := bleve.NewTextFieldMapping()
	userName.Store = false
	userName.Analyzer = options.PrimaryLanguage

	name.AddFieldMappingsAt("user_name", userName)

	annotation := bleve.NewTextFieldMapping()
	annotation.Store = false
	annotation.Analyzer = options.PrimaryLanguage

	name.AddFieldMappingsAt("annotation", annotation)

	memorizedFirstClipName := bleve.NewTextFieldMapping()
	memorizedFirstClipName.Store = false
	memorizedFirstClipName.Analyzer = options.PrimaryLanguage

	name.AddFieldMappingsAt("memorized_first_clip_name", memorizedFirstClipName)

	return name
}
