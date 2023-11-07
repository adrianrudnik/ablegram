package abletonv5

import (
	"github.com/adrianrudnik/ablegram/internal/pipeline"
	"github.com/adrianrudnik/ablegram/internal/stats"
	"github.com/adrianrudnik/ablegram/internal/tagger"
)

func ParsePreHearTracks(stat *stats.Statistics, path string, data *XmlRoot) []*pipeline.DocumentToIndexMsg {
	docs := make([]*pipeline.DocumentToIndexMsg, 0, 10)

	for _, returnTrack := range data.LiveSet.Tracks.PreHearTrack {
		tags := tagger.NewTagger()
		tags.Add("type:ableton-pre-hear-track")

		doc := NewPreHearTrackDocument()
		doc.LoadDisplayName([]string{
			returnTrack.Name.UserName.Value,
			returnTrack.Name.EffectiveName.Value,
		})
		doc.LoadFileReference(path, tags)
		doc.LoadTrackUserNames(&returnTrack.XmlTrackNameNode, tags)
		doc.LoadColor(returnTrack.Color.Value, tags)

		doc.EngraveTags(tags)

		docs = append(docs, pipeline.NewDocumentToIndexMsg(doc.GetAutoId(), doc))

		stat.IncrementCounter(AbletonPreHearTrack)
	}

	return docs
}
