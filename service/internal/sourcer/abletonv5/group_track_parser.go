package abletonv5

import (
	"github.com/adrianrudnik/ablegram/internal/pipeline"
	"github.com/adrianrudnik/ablegram/internal/stats"
	"github.com/adrianrudnik/ablegram/internal/tagger"
)

func ParseGroupTracks(stat *stats.Statistics, path string, data *XmlRoot) []*pipeline.DocumentToIndexMsg {
	docs := make([]*pipeline.DocumentToIndexMsg, 0, 10)

	for _, groupTrack := range data.LiveSet.Tracks.GroupTracks {
		tags := tagger.NewTagger()
		tags.Add("type:ableton-group-track")

		doc := NewReturnTrackDocument()
		doc.LoadDisplayName([]string{
			groupTrack.Name.UserName.Value,
			groupTrack.Name.EffectiveName.Value,
		})
		doc.LoadFileReference(path, tags)
		doc.LoadTrackUserNames(&groupTrack.Name, tags)
		doc.LoadColor(groupTrack.Color.Value, tags)

		doc.EngraveTags(tags)

		docs = append(docs, pipeline.NewDocumentToIndexMsg(doc.GetAutoId(), doc))

		stat.IncrementCounter(AbletonGroupTrack)
	}

	return docs
}
