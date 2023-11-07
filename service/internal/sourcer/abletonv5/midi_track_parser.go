package abletonv5

import (
	"github.com/adrianrudnik/ablegram/internal/pipeline"
	"github.com/adrianrudnik/ablegram/internal/stats"
	"github.com/adrianrudnik/ablegram/internal/tagger"
)

func ParseMidiTracks(
	stat *stats.Statistics,
	tc *tagger.TagCollector,
	path string,
	data *XmlRoot,
) []*pipeline.DocumentToIndexMsg {
	docs := make([]*pipeline.DocumentToIndexMsg, 0, 10)

	for _, midiTrack := range data.LiveSet.Tracks.MidiTracks {
		tb := tc.NewBucket()
		tb.Add("type:ableton-midi-track")

		doc := NewMidiTrackDocument()
		doc.LoadDisplayName([]string{
			midiTrack.Name.UserName.Value,
			midiTrack.Name.EffectiveName.Value,
		})
		doc.LoadFileReference(path, tb)
		doc.LoadTrackUserNames(&midiTrack.XmlTrackNameNode, tb)
		doc.LoadTrackIsFrozen(midiTrack.IsFrozen.Value, tb)
		doc.LoadColor(midiTrack.Color.Value, tb)

		doc.EngraveTags(tb)

		docs = append(docs, pipeline.NewDocumentToIndexMsg(doc.GetAutoId(), doc))

		stat.IncrementCounter(AbletonMidiTrack)
	}

	return docs
}
