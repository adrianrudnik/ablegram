package search

import (
	"github.com/blevesearch/bleve"
)

type Search struct {
	options *SearchOptions

	Index bleve.Index
}

func NewSearch(options *SearchOptions) (*Search, error) {
	indexMapping := bleve.NewIndexMapping()
	index, err := bleve.NewMemOnly(indexMapping)
	if err != nil {
		return nil, err
	}

	indexMapping.AddDocumentMapping("audio_track", createAudioTrackMapping(options))

	return &Search{
		options: options,
		Index:   index,
	}, nil
}
