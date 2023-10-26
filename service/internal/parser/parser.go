package parser

import (
	"encoding/xml"
	"github.com/adrianrudnik/ablegram/internal/parser/ablv5parser"
	"github.com/adrianrudnik/ablegram/internal/parser/ablv5schema"
	"github.com/adrianrudnik/ablegram/internal/pipeline"
	"github.com/adrianrudnik/ablegram/internal/stats"
	"github.com/rs/zerolog"
	"os"
)

var Logger = zerolog.New(os.Stderr).With().Timestamp().Logger()

func parseAlsV5(path string, m *stats.Metrics) ([]*pipeline.DocumentToIndexMsg, error) {
	rawContent, err := extractGzip(path)
	if err != nil {
		return nil, err
	}

	var data ablv5schema.Ableton

	err = xml.Unmarshal(rawContent, &data)
	if err != nil {
		return nil, err
	}

	// Create a slice to hold all documents that we out of the XML information
	docs := make([]*pipeline.DocumentToIndexMsg, 0, 50)

	docs = append(docs, ablv5parser.ParseLiveSet(m, path, &data))
	docs = append(docs, ablv5parser.ParseMidiTracks(m, path, &data)...)
	docs = append(docs, ablv5parser.ParseAudioTrack(m, path, &data)...)

	//for _, audioTrack := range data.LiveSet.Tracks.AudioTracks {
	//	id := fmt.Sprintf("%s_%s", path, audioTrack.Name.EffectiveName.Value)
	//	payload := indexer.AudioTrackDocument{
	//		Name: indexer.NameVariantDocument{
	//			DisplayName:            "bla",
	//			EffectiveName:          audioTrack.Name.EffectiveName.Value,
	//			UserName:               audioTrack.Name.UserName.Value,
	//			Annotation:             audioTrack.Name.Annotation.Value,
	//			MemorizedFirstClipName: audioTrack.Name.MemorizedFirstClipName.Value,
	//		},
	//		Filename: path,
	//	}
	//
	//	doc := pipeline.NewDocumentToIndexMsg(id, payload)
	//	docs = append(docs, doc)
	//
	//	m.AddAudioTrack()
	//}

	return docs, nil
}

func ParseAls(path string, m *stats.Metrics) ([]*pipeline.DocumentToIndexMsg, error) {
	Logger.Debug().Str("path", path).Msg("Start processing")

	r, err := parseAlsV5(path, m)
	if err != nil {
		m.AddInvalidFile()
		return nil, err
	}

	m.AddValidFile()
	return r, nil
}
