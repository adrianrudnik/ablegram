package abletonv5

import (
	"github.com/adrianrudnik/ablegram/internal/pipeline"
	"github.com/adrianrudnik/ablegram/internal/stats"
	"github.com/adrianrudnik/ablegram/internal/tagger"
)

func ParseAudioTracks(stat *stats.Statistics, path string, data *XmlRoot) []*pipeline.DocumentToIndexMsg {
	docs := make([]*pipeline.DocumentToIndexMsg, 0, 10)

	for _, audioTrack := range data.LiveSet.Tracks.AudioTracks {
		tags := tagger.NewTagger()
		tags.Add("type:ableton-audio-track")

		doc := NewAudioTrackDocument()
		doc.LoadDisplayName([]string{
			audioTrack.Name.UserName.Value,
			audioTrack.Name.EffectiveName.Value,
		})
		doc.LoadFileReference(path, tags)
		doc.LoadTrackUserNames(&audioTrack.Name, tags)
		doc.LoadColor(audioTrack.Color.Value, tags)

		doc.Frozen = audioTrack.Frozen.Value

		doc.EngraveTags(tags)

		docs = append(docs, pipeline.NewDocumentToIndexMsg(doc.GetAutoId(), doc))

		stat.IncrementCounter(AbletonAudioTrack)
	}

	return docs
}
