package abletonv5

import (
	"github.com/adrianrudnik/ablegram/internal/pipeline"
	"github.com/adrianrudnik/ablegram/internal/stats"
	"github.com/adrianrudnik/ablegram/internal/tagger"
)

func ParseReturnTracks(stat *stats.Statistics, path string, data *XmlRoot) []*pipeline.DocumentToIndexMsg {
	docs := make([]*pipeline.DocumentToIndexMsg, 0, 10)

	for _, returnTrack := range data.LiveSet.Tracks.ReturnTracks {
		tags := tagger.NewTagger()
		tags.AddSystemTag("type:ableton-return-track")

		doc := NewReturnTrackDocument()
		doc.LoadDisplayName([]string{
			returnTrack.Name.UserName.Value,
			returnTrack.Name.EffectiveName.Value,
		})
		doc.LoadFileReference(path, tags)
		doc.LoadTrackUserNames(&returnTrack.Name, tags)
		doc.LoadColor(returnTrack.Color.Value, tags)

		doc.EngraveTags(tags)

		docs = append(docs, pipeline.NewDocumentToIndexMsg(doc.GetAutoId(), doc))

		stat.IncrementCounter(AbletonReturnTrack)
	}

	return docs
}
