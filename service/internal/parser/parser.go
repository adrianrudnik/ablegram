package parser

import (
	"encoding/xml"
	"github.com/adrianrudnik/ablegram/internal/sourcer/abletonsrc"
	"github.com/adrianrudnik/ablegram/internal/stats"
	"github.com/adrianrudnik/ablegram/internal/tagger"
	"github.com/adrianrudnik/ablegram/internal/workload"
	"github.com/rs/zerolog"
	"os"
)

var Logger = zerolog.New(os.Stderr).With().Timestamp().Logger()

func parseAlsV5(
	stat *stats.Statistics,
	tc *tagger.TagCollector,
	path string,
) ([]*workload.DocumentPayload, error) {
	rawContent, err := extractGzip(path)
	if err != nil {
		return nil, err
	}

	var data abletonsrc.XmlRoot

	err = xml.Unmarshal(rawContent, &data)
	if err != nil {
		return nil, err
	}

	// Create a slice to hold all documents that we out of the XML information
	docs := make([]*workload.DocumentPayload, 0, 500)

	// Quality checked
	docs = append(docs, abletonsrc.ParseLiveSet(stat, tc, path, &data))

	// Not quality checked
	docs = append(docs, abletonsrc.ParseMidiTracks(stat, tc, path, &data)...)
	docs = append(docs, abletonsrc.ParseAudioTracks(stat, tc, path, &data)...)
	docs = append(docs, abletonsrc.ParseReturnTracks(stat, tc, path, &data)...)
	docs = append(docs, abletonsrc.ParseGroupTracks(stat, tc, path, &data)...)
	docs = append(docs, abletonsrc.ParsePreHearTracks(stat, tc, path, &data)...)
	docs = append(docs, abletonsrc.ParseMixerDocuments(stat, tc, path, &data)...)
	docs = append(docs, abletonsrc.ParseTrackDeviceChains(stat, tc, path, &data)...)
	docs = append(docs, abletonsrc.ParseScenes(stat, tc, path, &data)...)
	docs = append(docs, abletonsrc.ParseClips(stat, tc, path, &data)...)
	docs = append(docs, abletonsrc.ParseSampleReferences(stat, tc, path, &data)...)
	docs = append(docs, abletonsrc.ParseInfotext(stat, tc, path, &data)...)

	// Devices

	docs = append(docs, abletonsrc.ParseMidiArpeggiatorDevice(stat, tc, path, &data)...)
	docs = append(docs, abletonsrc.ParseMidiChordDevice(stat, tc, path, &data)...)
	docs = append(docs, abletonsrc.ParseMidiPitcherDevice(stat, tc, path, &data)...)
	docs = append(docs, abletonsrc.ParseMidiVelocityDevice(stat, tc, path, &data)...)

	// Finally create a file document, that bundles all found tags together
	// This allows us to search for files by tags, skipping detailed elements
	docs = append(docs, abletonsrc.ParseAlsFile(stat, tc, path, &data))

	return docs, nil
}

func ParseAls(
	stat *stats.Statistics,
	tc *tagger.TagCollector,
	path string,
) ([]*workload.DocumentPayload, error) {
	Logger.Debug().Str("path", path).Msg("Start processing")

	r, err := parseAlsV5(stat, tc, path)
	if err != nil {
		stat.IncrementCounter(stats.FileInvalid)
		return nil, err
	}

	stat.IncrementCounter(stats.FileValid)

	return r, nil
}
