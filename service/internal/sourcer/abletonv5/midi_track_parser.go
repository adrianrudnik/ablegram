package abletonv5

import (
	"github.com/adrianrudnik/ablegram/internal/pipeline"
	"github.com/adrianrudnik/ablegram/internal/stats"
	"github.com/adrianrudnik/ablegram/internal/tagger"
)

func ParseMidiTracks(stat *stats.Statistics, path string, data *XmlRoot) []*pipeline.DocumentToIndexMsg {
	docs := make([]*pipeline.DocumentToIndexMsg, 0, 10)

	for _, midiTrack := range data.LiveSet.Tracks.MidiTracks {
		tags := tagger.NewTagger()
		tags.Add("type:ableton-midi-track")

		doc := NewMidiTrackDocument()
		doc.LoadDisplayName([]string{
			midiTrack.Name.UserName.Value,
			midiTrack.Name.EffectiveName.Value,
		})
		doc.LoadFileReference(path, tags)
		doc.LoadTrackUserNames(&midiTrack.Name, tags)
		doc.LoadColor(midiTrack.Color.Value, tags)

		doc.Frozen = midiTrack.Frozen.Value

		doc.EngraveTags(tags)

		docs = append(docs, pipeline.NewDocumentToIndexMsg(doc.GetAutoId(), doc))

		stat.IncrementCounter(AbletonMidiTrack)
	}

	return docs
}
