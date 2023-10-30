package indexer

import (
	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/analysis/analyzer/custom"
	"github.com/blevesearch/bleve/v2/analysis/lang/en"
	"github.com/blevesearch/bleve/v2/analysis/token/edgengram"
	"github.com/blevesearch/bleve/v2/analysis/token/lowercase"
	"github.com/blevesearch/bleve/v2/analysis/tokenizer/unicode"
	"github.com/rs/zerolog"
	"os"
)

var Logger = zerolog.New(os.Stderr).With().Timestamp().Logger()

type Search struct {
	Index bleve.Index
}

func NewSearch() *Search {
	indexMapping := bleve.NewIndexMapping()

	index, err := bleve.NewMemOnly(indexMapping)
	if err != nil {
		Logger.Panic().Err(err).Msg("Failed to create memory index")
		panic(err)
	}

	err = indexMapping.AddCustomTokenFilter("edgeNgram325",
		map[string]interface{}{
			"type": edgengram.Name,
			"min":  3.0,
			"max":  25.0,
		})
	if err != nil {
		Logger.Panic().Err(err).Msg("Failed to register edgeNgram325 token filter")
		panic(err)
	}

	err = indexMapping.AddCustomAnalyzer("enWithEdgeNgram325",
		map[string]interface{}{
			"type":      custom.Name,
			"tokenizer": unicode.Name,
			"token_filters": []string{
				en.PossessiveName,
				lowercase.Name,
				en.StopName,
				"edgeNgram325",
			},
		})
	if err != nil {
		Logger.Panic().Err(err).Msg("Failed to register enWithEdgeNgram325 custom analyzer")
		panic(err)
	}

	indexMapping.DefaultAnalyzer = en.AnalyzerName

	indexMapping.AddDocumentMapping("AudioTrack", buildAudioTrackMapping())
	indexMapping.AddDocumentMapping("LiveSet", buildLiveSetMapping())
	indexMapping.AddDocumentMapping("MidiTrack", buildMidiTrackMapping())

	Logger.Info().Msg("Index documents mapped")

	return &Search{
		Index: index,
	}
}
