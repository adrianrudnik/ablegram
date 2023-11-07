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

func parseAlsV5(stat *stats.Statistics, path string) ([]*pipeline.DocumentToIndexMsg, error) {
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
	docs := make([]*pipeline.DocumentToIndexMsg, 0, 200)

	docs = append(docs, abletonv5.ParseLiveSet(stat, path, &data))
	docs = append(docs, abletonv5.ParseMidiTracks(stat, path, &data)...)
	docs = append(docs, abletonv5.ParseAudioTracks(stat, path, &data)...)
	docs = append(docs, abletonv5.ParseReturnTracks(stat, path, &data)...)
	docs = append(docs, abletonv5.ParseGroupTracks(stat, path, &data)...)
	docs = append(docs, abletonv5.ParsePreHearTracks(stat, path, &data)...)
	docs = append(docs, abletonv5.ParseMixerDocuments(stat, path, &data)...)
	docs = append(docs, abletonv5.ParseTrackDeviceChains(stat, path, &data)...)
	docs = append(docs, abletonv5.ParseScenes(stat, path, &data)...)
	docs = append(docs, abletonv5.ParseClips(stat, path, &data)...)

	// Devices

	docs = append(docs, abletonv5.ParseMidiArpeggiatorDevice(stat, path, &data)...)
	docs = append(docs, abletonv5.ParseMidiChordDevice(stat, path, &data)...)

	return docs, nil
}

func ParseAls(stat *stats.Statistics, path string) ([]*pipeline.DocumentToIndexMsg, error) {
	Logger.Debug().Str("path", path).Msg("Start processing")

	r, err := parseAlsV5(stat, path)
	if err != nil {
		stat.IncrementCounter(stats.FileInvalid)
		return nil, err
	}

	stat.IncrementCounter(stats.FileValid)

	return r, nil
}
