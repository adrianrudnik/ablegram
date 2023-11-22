package abletonsrc

import (
	"github.com/adrianrudnik/ablegram/internal/stats"
	"github.com/adrianrudnik/ablegram/internal/tagger"
	"github.com/adrianrudnik/ablegram/internal/workload"
)

func ParseMixerDocuments(
	stat *stats.Statistics,
	tc *tagger.TagCollector,
	path string,
	data *XmlRoot,
) []*workload.DocumentPayload {
	docs := make([]*workload.DocumentPayload, 0, 10)

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
		tb := tc.NewBucket()
		tb.Add("type:ableton-mixer")

		doc := NewMixerDocument()
		doc.LoadDisplayName([]string{})
		doc.LoadFileReference(path, tb)
		doc.LoadUserInfoText(mx.Annotation.Value, tb)
		doc.LoadUserName(mx.UserName.Value, tb)
		doc.LoadOptionFolded(mx.IsFolded.Value, tb)
		doc.LoadOptionExpanded(mx.IsExpanded.Value, tb)

		doc.EngraveTags(tb)

		docs = append(docs, workload.NewDocumentPayload(doc.GetAutoId(), doc))

		stat.IncrementCounter(AbletonMixer)
	}

	return docs
}
