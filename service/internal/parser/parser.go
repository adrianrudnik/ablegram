package parser

import (
	"encoding/xml"
	"fmt"
	"github.com/adrianrudnik/ablegram/internal/parser/ablv5schema"
	"github.com/adrianrudnik/ablegram/internal/pipeline"
	search2 "github.com/adrianrudnik/ablegram/internal/search"
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

	docs := make([]*pipeline.DocumentToIndexMsg, 0, 10)

	for _, midiTrack := range data.LiveSet.Tracks.MidiTracks {
		id := fmt.Sprintf("%s_%s", path, midiTrack.Name.EffectiveName.Value)
		payload := search2.MidiTrackDocument{
			Name: search2.NameVariantDocument{
				DisplayName:            "bla",
				EffectiveName:          midiTrack.Name.EffectiveName.Value,
				UserName:               midiTrack.Name.UserName.Value,
				Annotation:             midiTrack.Name.Annotation.Value,
				MemorizedFirstClipName: midiTrack.Name.MemorizedFirstClipName.Value,
			},
			Filename: path,
		}

		doc := pipeline.NewDocumentToIndexMsg(id, payload)
		docs = append(docs, doc)

		m.AddMidiTrack()
	}

	for _, audioTrack := range data.LiveSet.Tracks.AudioTracks {
		id := fmt.Sprintf("%s_%s", path, audioTrack.Name.EffectiveName.Value)
		payload := search2.AudioTrackDocument{
			Name: search2.NameVariantDocument{
				DisplayName:            "bla",
				EffectiveName:          audioTrack.Name.EffectiveName.Value,
				UserName:               audioTrack.Name.UserName.Value,
				Annotation:             audioTrack.Name.Annotation.Value,
				MemorizedFirstClipName: audioTrack.Name.MemorizedFirstClipName.Value,
			},
			Filename: path,
		}

		doc := pipeline.NewDocumentToIndexMsg(id, payload)
		docs = append(docs, doc)

		m.AddAudioTrack()
	}

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
