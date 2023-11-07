package abletonv5

import (
	"github.com/adrianrudnik/ablegram/internal/pipeline"
	"github.com/adrianrudnik/ablegram/internal/stats"
	"github.com/adrianrudnik/ablegram/internal/tagger"
)

func ParseGroupTracks(
	stat *stats.Statistics,
	tc *tagger.TagCollector,
	path string,
	data *XmlRoot,
) []*pipeline.DocumentToIndexMsg {
	docs := make([]*pipeline.DocumentToIndexMsg, 0, 10)

	for _, groupTrack := range data.LiveSet.Tracks.GroupTracks {
		tb := tc.NewBucket()
		tb.Add("type:ableton-group-track")

		doc := NewReturnTrackDocument()
		doc.LoadDisplayName([]string{
			groupTrack.Name.UserName.Value,
			groupTrack.Name.EffectiveName.Value,
		})
		doc.LoadFileReference(path, tb)
		doc.LoadTrackUserNames(&groupTrack.XmlTrackNameNode, tb)
		doc.LoadColor(groupTrack.Color.Value, tb)

		doc.EngraveTags(tb)

		docs = append(docs, pipeline.NewDocumentToIndexMsg(doc.GetAutoId(), doc))

		stat.IncrementCounter(AbletonGroupTrack)
	}

	return docs
}
