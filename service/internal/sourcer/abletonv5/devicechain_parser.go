package abletonv5

import (
	"github.com/adrianrudnik/ablegram/internal/pipeline"
	"github.com/adrianrudnik/ablegram/internal/stats"
	"github.com/adrianrudnik/ablegram/internal/tagger"
)

func ParseDeviceChains(stat *stats.Statistics, path string, data *XmlRoot) []*pipeline.DocumentToIndexMsg {
	docs := make([]*pipeline.DocumentToIndexMsg, 0, 10)

	found := make([]XmlDeviceChain, 0, 100)

	// Collect all possible device chains
	for _, midiTrack := range data.LiveSet.Tracks.MidiTracks {
		found = append(found, midiTrack.DeviceChain)
	}

	for _, audioTrack := range data.LiveSet.Tracks.AudioTracks {
		found = append(found, audioTrack.DeviceChain)
	}

	for _, groupTrack := range data.LiveSet.Tracks.GroupTracks {
		found = append(found, groupTrack.DeviceChain)
	}

	for _, returnTrack := range data.LiveSet.Tracks.ReturnTracks {
		found = append(found, returnTrack.DeviceChain)
	}

	for _, dc := range found {
		tags := tagger.NewTagger()
		tags.AddSystemTag("type:ableton-device-chain")

		doc := NewDeviceChainDocument()
		doc.LoadDisplayName([]string{AbletonDeviceChain})
		doc.LoadFileReference(path, tags)

		doc.DeviceCount = dc.Devices.GetCount()

		doc.EngraveTags(tags)

		docs = append(docs, pipeline.NewDocumentToIndexMsg(doc.GetAutoId(), doc))

		stat.IncrementCounter(AbletonDeviceChain)
	}

	return docs
}
