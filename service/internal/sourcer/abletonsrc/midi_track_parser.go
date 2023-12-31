package abletonsrc

import (
	"github.com/adrianrudnik/ablegram/internal/stats"
	"github.com/adrianrudnik/ablegram/internal/tagger"
	"github.com/adrianrudnik/ablegram/internal/workload"
)

func ParseMidiTracks(
	stat *stats.Statistics,
	tc *tagger.TagCollector,
	path string,
	data *XmlRoot,
) []*workload.DocumentPayload {
	docs := make([]*workload.DocumentPayload, 0, 10)

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
		doc.LoadIsFrozenOption(midiTrack.IsFrozen.Value, tb)
		doc.LoadColor(midiTrack.Color.Value, tb)

		doc.EngraveTags(tb)

		docs = append(docs, workload.NewDocumentPayload(doc.GetAutoId(), doc))

		stat.IncrementCounter(AbletonMidiTrack)
	}

	return docs
}
