package ablv5parser

import (
	"fmt"
	"github.com/adrianrudnik/ablegram/internal/indexer"
	"github.com/adrianrudnik/ablegram/internal/parser/ablv5schema"
	"github.com/adrianrudnik/ablegram/internal/pipeline"
	"github.com/adrianrudnik/ablegram/internal/stats"
	"github.com/adrianrudnik/ablegram/internal/tagger"
	"github.com/adrianrudnik/ablegram/internal/util"
	"path/filepath"
)

func ParseAudioTrack(m *stats.Metrics, path string, data *ablv5schema.Ableton) []*pipeline.DocumentToIndexMsg {
	docs := make([]*pipeline.DocumentToIndexMsg, 0, 10)

	for _, audioTrack := range data.LiveSet.Tracks.AudioTracks {
		tags := tagger.NewTagger()
		tags.AddSystemTag("type:audio-track")

		// Derive document
		id := tagger.IdHash(fmt.Sprintf("%s_%s", path, audioTrack.Name.EffectiveName.Value))

		displayName := []string{
			audioTrack.Name.UserName.Value,
			audioTrack.Name.EffectiveName.Value,
		}

		doc := indexer.NewAudioTrackDocument()

		doc.PathAbsolute = path
		doc.PathFolder = filepath.Dir(path)
		doc.Filename = filepath.Base(path)

		doc.DisplayName = util.Namelize(displayName)
		doc.EffectiveName = audioTrack.Name.EffectiveName.Value
		doc.UserName = audioTrack.Name.UserName.Value
		doc.MemorizedFirstClipName = audioTrack.Name.MemorizedFirstClipName.Value

		doc.Color = audioTrack.Color.Value
		tags.AddSystemTag(fmt.Sprintf("color:all:%d", audioTrack.Color.Value))
		tags.AddSystemTag(fmt.Sprintf("color:track:%d", audioTrack.Color.Value))

		// Annotation
		val, empty := util.EvaluateUserInput(audioTrack.Name.Annotation.Value)
		if !empty {
			doc.Annotation = val
			tags.AddSystemTag("info:has-annotation")
		} else {
			tags.AddSystemTag("info:no-annotation")
		}

		doc.Tags = tags.GetAllAndClear()

		docs = append(docs, pipeline.NewDocumentToIndexMsg(id, doc))

		m.CountAudioTrack()
	}

	return docs
}
