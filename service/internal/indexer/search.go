package indexer

import (
	"github.com/adrianrudnik/ablegram/internal/sourcer/abletonsrc"
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

	err = indexMapping.AddCustomTokenFilter("edgeNgram225",
		map[string]interface{}{
			"type": edgengram.Name,
			"min":  2.0,
			"max":  25.0,
		})
	if err != nil {
		Logger.Panic().Err(err).Msg("Failed to register edgeNgram225 token filter")
		panic(err)
	}

	err = indexMapping.AddCustomAnalyzer("enWithEdgeNgram225",
		map[string]interface{}{
			"type":      custom.Name,
			"tokenizer": unicode.Name,
			"token_filters": []string{
				en.PossessiveName,
				lowercase.Name,
				en.StopName,
				"edgeNgram225",
			},
		})
	if err != nil {
		Logger.Panic().Err(err).Msg("Failed to register enWithEdgeNgram225 custom analyzer")
		panic(err)
	}

	indexMapping.DefaultAnalyzer = en.AnalyzerName

	abletonsrc.RegisterDocumentMappings(indexMapping)

	Logger.Info().Msg("Index documents mapped")

	return &Search{
		Index: index,
	}
}
