package abletonv5

import (
	"fmt"
	"github.com/adrianrudnik/ablegram/internal/pipeline"
	"github.com/adrianrudnik/ablegram/internal/stats"
	"github.com/adrianrudnik/ablegram/internal/tagger"
	"github.com/adrianrudnik/ablegram/internal/util"
	"path/filepath"
)

func ParseAudioTracks(m *stats.Metrics, path string, data *Ableton) []*pipeline.DocumentToIndexMsg {
	docs := make([]*pipeline.DocumentToIndexMsg, 0, 10)

	for _, audioTrack := range data.LiveSet.Tracks.AudioTracks {
		tags := tagger.NewTagger()
		tags.AddSystemTag("type:ableton-audio-track")

		// Derive document
		id := tagger.IdHash(fmt.Sprintf("%s_%s", path, audioTrack.Name.EffectiveName.Value))

		displayName := []string{
			audioTrack.Name.UserName.Value,
			audioTrack.Name.EffectiveName.Value,
		}

		doc := NewAudioTrackDocument()

		doc.PathAbsolute = path
		doc.PathFolder = filepath.Dir(path)
		doc.Filename = filepath.Base(path)

		doc.DisplayName = util.Namelize(displayName)
		doc.EffectiveName = audioTrack.Name.EffectiveName.Value
		doc.UserName = audioTrack.Name.UserName.Value
		doc.MemorizedFirstClipName = audioTrack.Name.MemorizedFirstClipName.Value
		doc.Annotation = parseAnnotation(tags, audioTrack.Name.Annotation.Value)

		doc.Color = parseColor(tags, audioTrack.Color.Value)
		doc.Frozen = audioTrack.Frozen.Value

		doc.Tags = tags.GetAllAndClear()

		docs = append(docs, pipeline.NewDocumentToIndexMsg(id, doc))

		m.CountAudioTrack()
	}

	return docs
}
