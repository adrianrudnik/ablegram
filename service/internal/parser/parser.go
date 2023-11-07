package parser

import (
	"encoding/xml"
	"github.com/adrianrudnik/ablegram/internal/pipeline"
	"github.com/adrianrudnik/ablegram/internal/sourcer/abletonv5"
	"github.com/adrianrudnik/ablegram/internal/stats"
	"github.com/adrianrudnik/ablegram/internal/tagger"
	"github.com/rs/zerolog"
	"os"
)

var Logger = zerolog.New(os.Stderr).With().Timestamp().Logger()

func parseAlsV5(stat *stats.Statistics, tc *tagger.TagCollector, path string) ([]*pipeline.DocumentToIndexMsg, error) {
	rawContent, err := extractGzip(path)
	if err != nil {
		return nil, err
	}

	var data abletonv5.XmlRoot

	err = xml.Unmarshal(rawContent, &data)
	if err != nil {
		return nil, err
	}

	// Create a slice to hold all documents that we out of the XML information
	docs := make([]*pipeline.DocumentToIndexMsg, 0, 500)

	docs = append(docs, abletonv5.ParseLiveSet(stat, tc, path, &data))
	docs = append(docs, abletonv5.ParseMidiTracks(stat, tc, path, &data)...)
	docs = append(docs, abletonv5.ParseAudioTracks(stat, tc, path, &data)...)
	docs = append(docs, abletonv5.ParseReturnTracks(stat, tc, path, &data)...)
	docs = append(docs, abletonv5.ParseGroupTracks(stat, tc, path, &data)...)
	docs = append(docs, abletonv5.ParsePreHearTracks(stat, tc, path, &data)...)
	docs = append(docs, abletonv5.ParseMixerDocuments(stat, tc, path, &data)...)
	docs = append(docs, abletonv5.ParseTrackDeviceChains(stat, tc, path, &data)...)
	docs = append(docs, abletonv5.ParseScenes(stat, tc, path, &data)...)
	docs = append(docs, abletonv5.ParseClips(stat, tc, path, &data)...)

	// Devices

	docs = append(docs, abletonv5.ParseMidiArpeggiatorDevice(stat, tc, path, &data)...)
	docs = append(docs, abletonv5.ParseMidiChordDevice(stat, tc, path, &data)...)
	docs = append(docs, abletonv5.ParseMidiPitcherDevice(stat, tc, path, &data)...)
	docs = append(docs, abletonv5.ParseMidiVelocityDevice(stat, tc, path, &data)...)

	return docs, nil
}

func ParseAls(
	stat *stats.Statistics,
	tags *tagger.TagCollector,
	path string,
) ([]*pipeline.DocumentToIndexMsg, error) {
	Logger.Debug().Str("path", path).Msg("Start processing")

	r, err := parseAlsV5(stat, tags, path)
	if err != nil {
		stat.IncrementCounter(stats.FileInvalid)
		return nil, err
	}

	stat.IncrementCounter(stats.FileValid)

	return r, nil
}
