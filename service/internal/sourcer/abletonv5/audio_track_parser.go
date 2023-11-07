package abletonv5

import (
	"github.com/adrianrudnik/ablegram/internal/pipeline"
	"github.com/adrianrudnik/ablegram/internal/stats"
	"github.com/adrianrudnik/ablegram/internal/tagger"
)

func ParseAudioTracks(
	stat *stats.Statistics,
	tc *tagger.TagCollector,
	path string,
	data *XmlRoot,
) []*pipeline.DocumentToIndexMsg {
	docs := make([]*pipeline.DocumentToIndexMsg, 0, 10)

	for _, audioTrack := range data.LiveSet.Tracks.AudioTracks {
		tb := tc.NewBucket()
		tb.Add("type:ableton-audio-track")

		doc := NewAudioTrackDocument()
		doc.LoadDisplayName([]string{
			audioTrack.Name.UserName.Value,
			audioTrack.Name.EffectiveName.Value,
		})
		doc.LoadFileReference(path, tb)
		doc.LoadTrackUserNames(&audioTrack.XmlTrackNameNode, tb)
		doc.LoadTrackIsFrozen(audioTrack.IsFrozen.Value, tb)
		doc.LoadColor(audioTrack.Color.Value, tb)

		doc.EngraveTags(tb)

		docs = append(docs, pipeline.NewDocumentToIndexMsg(doc.GetAutoId(), doc))

		stat.IncrementCounter(AbletonAudioTrack)
	}

	return docs
}
