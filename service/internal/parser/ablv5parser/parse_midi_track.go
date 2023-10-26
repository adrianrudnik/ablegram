package ablv5parser

import (
	"fmt"
	"github.com/adrianrudnik/ablegram/internal/indexer"
	"github.com/adrianrudnik/ablegram/internal/parser/ablv5schema"
	"github.com/adrianrudnik/ablegram/internal/pipeline"
	"github.com/adrianrudnik/ablegram/internal/stats"
	"github.com/adrianrudnik/ablegram/internal/tagger"
	"github.com/adrianrudnik/ablegram/internal/util"
)

func ParseMidiTracks(m *stats.Metrics, path string, data *ablv5schema.Ableton) []*pipeline.DocumentToIndexMsg {
	tags := tagger.NewTagger()

	docs := make([]*pipeline.DocumentToIndexMsg, 0, 10)

	for _, midiTrack := range data.LiveSet.Tracks.MidiTracks {
		id := fmt.Sprintf("%s_%s", path, midiTrack.Name.EffectiveName.Value)

		displayName := []string{
			midiTrack.Name.UserName.Value,
			midiTrack.Name.EffectiveName.Value,
		}

		track := indexer.NewMidiTrackDocument()
		track.Tags = tags.GetAllAndClear()
		track.DisplayName = util.Namelize(displayName)
		track.EffectiveName = midiTrack.Name.EffectiveName.Value
		track.UserName = midiTrack.Name.UserName.Value
		track.Annotation = midiTrack.Name.Annotation.Value
		track.MemorizedFirstClipName = midiTrack.Name.MemorizedFirstClipName.Value
		track.Filename = path

		docs = append(docs, pipeline.NewDocumentToIndexMsg(id, track))

		m.AddMidiTrack()
	}

	return docs
}
