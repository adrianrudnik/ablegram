package abletonsrc

import (
	"github.com/adrianrudnik/ablegram/internal/stats"
	"github.com/adrianrudnik/ablegram/internal/tagger"
	"github.com/adrianrudnik/ablegram/internal/workload"
)

func ParseReturnTracks(
	stat *stats.Statistics,
	tc *tagger.TagCollector,
	path string,
	data *XmlRoot,
) []*workload.DocumentPayload {
	docs := make([]*workload.DocumentPayload, 0, 10)

	for _, returnTrack := range data.LiveSet.Tracks.ReturnTracks {
		tb := tc.NewBucket()
		tb.Add("type:ableton-return-track")

		doc := NewReturnTrackDocument()
		doc.LoadDisplayName([]string{
			returnTrack.Name.UserName.Value,
			returnTrack.Name.EffectiveName.Value,
		})
		doc.LoadFileReference(path, tb)
		doc.LoadTrackUserNames(&returnTrack.XmlTrackNameNode, tb)
		doc.LoadColor(returnTrack.Color.Value, tb)

		doc.EngraveTags(tb)

		docs = append(docs, workload.NewDocumentPayload(doc.GetAutoId(), doc))

		stat.IncrementCounter(AbletonReturnTrack)
	}

	return docs
}
