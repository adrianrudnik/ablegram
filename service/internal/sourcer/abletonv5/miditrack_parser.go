package abletonv5

import (
	"fmt"
	"github.com/adrianrudnik/ablegram/internal/pipeline"
	"github.com/adrianrudnik/ablegram/internal/stats"
	"github.com/adrianrudnik/ablegram/internal/tagger"
	"github.com/adrianrudnik/ablegram/internal/util"
	"path/filepath"
)

func ParseMidiTracks(m *stats.Metrics, path string, data *Ableton) []*pipeline.DocumentToIndexMsg {
	docs := make([]*pipeline.DocumentToIndexMsg, 0, 10)

	for _, midiTrack := range data.LiveSet.Tracks.MidiTracks {
		tags := tagger.NewTagger()
		tags.AddSystemTag("type:ableton-midi-track")

		// Derive document
		id := tagger.IdHash(fmt.Sprintf("%s_%s", path, midiTrack.Name.EffectiveName.Value))

		displayName := []string{
			midiTrack.Name.UserName.Value,
			midiTrack.Name.EffectiveName.Value,
		}

		doc := NewMidiTrackDocument()

		doc.PathAbsolute = path
		doc.PathFolder = filepath.Dir(path)
		doc.Filename = filepath.Base(path)

		doc.DisplayName = util.Namelize(displayName)
		doc.EffectiveName = midiTrack.Name.EffectiveName.Value
		doc.UserName = midiTrack.Name.UserName.Value
		doc.MemorizedFirstClipName = midiTrack.Name.MemorizedFirstClipName.Value
		doc.Annotation = parseAnnotation(tags, midiTrack.Name.Annotation.Value)

		doc.Color = parseColor(tags, midiTrack.Color.Value)
		doc.Frozen = midiTrack.Frozen.Value

		doc.Tags = tags.GetAllAndClear()

		docs = append(docs, pipeline.NewDocumentToIndexMsg(id, doc))

		m.CountMidiTrack()
	}

	return docs
}
