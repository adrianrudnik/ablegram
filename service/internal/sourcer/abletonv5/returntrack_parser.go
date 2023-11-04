package abletonv5

import (
	"fmt"
	"github.com/adrianrudnik/ablegram/internal/pipeline"
	"github.com/adrianrudnik/ablegram/internal/stats"
	"github.com/adrianrudnik/ablegram/internal/tagger"
	"github.com/adrianrudnik/ablegram/internal/util"
	"path/filepath"
)

func ParseReturnTracks(m *stats.Metrics, path string, data *Ableton) []*pipeline.DocumentToIndexMsg {
	docs := make([]*pipeline.DocumentToIndexMsg, 0, 10)

	for _, returnTrack := range data.LiveSet.Tracks.ReturnTracks {
		tags := tagger.NewTagger()
		tags.AddSystemTag("type:ableton-return-track")

		// Derive document
		id := tagger.IdHash(fmt.Sprintf("%s_%s", path, returnTrack.Name.EffectiveName.Value))

		displayName := []string{
			returnTrack.Name.UserName.Value,
			returnTrack.Name.EffectiveName.Value,
		}

		doc := NewReturnTrackDocument()

		doc.PathAbsolute = path
		doc.PathFolder = filepath.Dir(path)
		doc.Filename = filepath.Base(path)

		doc.DisplayName = util.Namelize(displayName)
		doc.EffectiveName = returnTrack.Name.EffectiveName.Value
		doc.UserName = returnTrack.Name.UserName.Value
		doc.MemorizedFirstClipName = returnTrack.Name.MemorizedFirstClipName.Value
		doc.Annotation = parseAnnotation(tags, returnTrack.Name.Annotation.Value)

		doc.Color = parseColor(tags, returnTrack.Color.Value)

		doc.Tags = tags.GetAllAndClear()

		docs = append(docs, pipeline.NewDocumentToIndexMsg(id, doc))

		m.CountAudioTrack()
	}

	return docs
}
