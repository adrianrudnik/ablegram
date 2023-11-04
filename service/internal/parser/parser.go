package parser

import (
	"encoding/xml"
	"github.com/adrianrudnik/ablegram/internal/pipeline"
	"github.com/adrianrudnik/ablegram/internal/sourcer/abletonv5"
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

	var data abletonv5.Ableton

	err = xml.Unmarshal(rawContent, &data)
	if err != nil {
		return nil, err
	}

	// Create a slice to hold all documents that we out of the XML information
	docs := make([]*pipeline.DocumentToIndexMsg, 0, 50)

	docs = append(docs, abletonv5.ParseLiveSet(m, path, &data))
	docs = append(docs, abletonv5.ParseMidiTracks(m, path, &data)...)
	docs = append(docs, abletonv5.ParseAudioTracks(m, path, &data)...)

	return docs, nil
}

func ParseAls(path string, m *stats.Metrics) ([]*pipeline.DocumentToIndexMsg, error) {
	Logger.Debug().Str("path", path).Msg("Start processing")

	r, err := parseAlsV5(path, m)
	if err != nil {
		go m.CountInvalidFile()
		return nil, err
	}

	m.CountValidFile()
	return r, nil
}
