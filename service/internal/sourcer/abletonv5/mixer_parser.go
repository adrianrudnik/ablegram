package abletonv5

import (
	"github.com/adrianrudnik/ablegram/internal/pipeline"
	"github.com/adrianrudnik/ablegram/internal/stats"
	"github.com/adrianrudnik/ablegram/internal/tagger"
)

func ParseMixerDocuments(stat *stats.Statistics, path string, data *XmlRoot) []*pipeline.DocumentToIndexMsg {
	docs := make([]*pipeline.DocumentToIndexMsg, 0, 10)

	found := make([]XmlMixer, 0, 100)

	// Collect all possible mixer tracks
	for _, midiTrack := range data.LiveSet.Tracks.MidiTracks {
		found = append(found, midiTrack.DeviceChain.Mixer)
	}

	for _, audioTrack := range data.LiveSet.Tracks.AudioTracks {
		found = append(found, audioTrack.DeviceChain.Mixer)
	}

	for _, groupTrack := range data.LiveSet.Tracks.GroupTracks {
		found = append(found, groupTrack.DeviceChain.Mixer)
	}

	for _, returnTrack := range data.LiveSet.Tracks.ReturnTracks {
		found = append(found, returnTrack.DeviceChain.Mixer)
	}

	for _, mx := range found {
		tags := tagger.NewTagger()
		tags.Add("type:ableton-mixer")

		doc := NewMixerDocument()
		doc.LoadDisplayName([]string{AbletonMixer})
		doc.LoadFileReference(path, tags)
		doc.LoadUserInfoText(mx.Annotation.Value, tags)
		doc.LoadUserName(mx.UserName.Value, tags)

		doc.EngraveTags(tags)

		docs = append(docs, pipeline.NewDocumentToIndexMsg(doc.GetAutoId(), doc))

		stat.IncrementCounter(AbletonMixer)
	}

	return docs
}
