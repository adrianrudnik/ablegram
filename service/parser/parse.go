package parser

import (
	"encoding/xml"
	"fmt"
	"github.com/adrianrudnik/ablegram/parser/ablv5schema"
	"github.com/adrianrudnik/ablegram/pipeline"
	"github.com/adrianrudnik/ablegram/search"
	"github.com/rs/zerolog"
	"os"
)

var Logger = zerolog.New(os.Stderr).With().Timestamp().Logger()

func parseAlsV5(path string) ([]*pipeline.DocumentToIndexMsg, error) {
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
		payload := search.MidiTrackDocument{
			Name: search.NameVariantDocument{
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
	}

	return docs, nil
}

func ParseAls(path string) ([]*pipeline.DocumentToIndexMsg, error) {
	Logger.Debug().Str("path", path).Msg("Start processing")
	r, err := parseAlsV5(path)
	if err != nil {
		return nil, err
	}

	return r, nil
}
