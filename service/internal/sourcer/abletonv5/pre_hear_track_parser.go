package abletonv5

import (
	"github.com/adrianrudnik/ablegram/internal/pipeline"
	"github.com/adrianrudnik/ablegram/internal/stats"
	"github.com/adrianrudnik/ablegram/internal/tagger"
)

func ParsePreHearTracks(
	stat *stats.Statistics,
	tc *tagger.TagCollector,
	path string,
	data *XmlRoot,
) []*pipeline.DocumentToIndexMsg {
	docs := make([]*pipeline.DocumentToIndexMsg, 0, 10)

	for _, returnTrack := range data.LiveSet.Tracks.PreHearTrack {
		tb := tc.NewBucket()
		tb.Add("type:ableton-pre-hear-track")

		doc := NewPreHearTrackDocument()
		doc.LoadDisplayName([]string{
			returnTrack.Name.UserName.Value,
			returnTrack.Name.EffectiveName.Value,
		})
		doc.LoadFileReference(path, tb)
		doc.LoadTrackUserNames(&returnTrack.XmlTrackNameNode, tb)
		doc.LoadColor(returnTrack.Color.Value, tb)

		doc.EngraveTags(tb)

		docs = append(docs, pipeline.NewDocumentToIndexMsg(doc.GetAutoId(), doc))

		stat.IncrementCounter(AbletonPreHearTrack)
	}

	return docs
}
