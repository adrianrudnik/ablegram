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

func ParseMidiTracks(m *stats.Metrics, path string, data *ablv5schema.Ableton) []*pipeline.DocumentToIndexMsg {
	docs := make([]*pipeline.DocumentToIndexMsg, 0, 10)

	for _, midiTrack := range data.LiveSet.Tracks.MidiTracks {
		tags := tagger.NewTagger()
		tags.AddSystemTag("type:midi-track")

		// Derive document
		id := tagger.IdHash(fmt.Sprintf("%s_%s", path, midiTrack.Name.EffectiveName.Value))

		displayName := []string{
			midiTrack.Name.UserName.Value,
			midiTrack.Name.EffectiveName.Value,
		}

		doc := indexer.NewMidiTrackDocument()

		doc.PathAbsolute = path
		doc.PathFolder = filepath.Dir(path)
		doc.Filename = filepath.Base(path)

		doc.DisplayName = util.Namelize(displayName)
		doc.EffectiveName = midiTrack.Name.EffectiveName.Value
		doc.UserName = midiTrack.Name.UserName.Value
		doc.MemorizedFirstClipName = midiTrack.Name.MemorizedFirstClipName.Value

		doc.Color = midiTrack.Color.Value
		tags.AddSystemTag(fmt.Sprintf("color:all:%d", midiTrack.Color.Value))
		tags.AddSystemTag(fmt.Sprintf("color:track:%d", midiTrack.Color.Value))

		// Annotation
		val, empty := util.EvaluateUserInput(midiTrack.Name.Annotation.Value)
		if !empty {
			doc.Annotation = val
			tags.AddSystemTag("info:has-annotation")
		} else {
			tags.AddSystemTag("info:no-annotation")
		}

		doc.Tags = tags.GetAllAndClear()

		docs = append(docs, pipeline.NewDocumentToIndexMsg(id, doc))

		m.CountMidiTrack()
	}

	return docs
}
