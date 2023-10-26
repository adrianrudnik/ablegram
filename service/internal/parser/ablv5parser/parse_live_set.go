package ablv5parser

import (
	"fmt"
	"github.com/adrianrudnik/ablegram/internal/indexer"
	"github.com/adrianrudnik/ablegram/internal/parser/ablv5schema"
	"github.com/adrianrudnik/ablegram/internal/pipeline"
	"github.com/adrianrudnik/ablegram/internal/stats"
	"github.com/adrianrudnik/ablegram/internal/tagger"
	"path/filepath"
	"strings"
)

func ParseLiveSet(m *stats.Metrics, path string, data *ablv5schema.Ableton) *pipeline.DocumentToIndexMsg {
	tags := tagger.NewTagger()

	if len(data.LiveSet.Tracks.AudioTracks) > 0 {
		tags.AddSystemTag("live-set:has-audio-track")
	} else {
		tags.AddSystemTag("live-set:no-audio-track")
	}

	if len(data.LiveSet.Tracks.MidiTracks) > 0 {
		tags.AddSystemTag("live-set:has-midi-track")
	} else {
		tags.AddSystemTag("live-set:no-midi-track")
	}

	if strings.HasPrefix(data.Creator, "Ableton Live 11 ") {
		tags.AddSystemTag(fmt.Sprintf("ableton:version:%s", strings.TrimPrefix(data.Creator, "Ableton Live ")))
	}

	liveSet := indexer.NewLiveSetDocument()
	liveSet.Tags = tags.GetAllAndClear()
	liveSet.DisplayName = filepath.Base(path)
	liveSet.Filename = path
	liveSet.MajorVersion = data.MajorVersion
	liveSet.MinorVersion = data.MinorVersion
	liveSet.Creator = data.Creator
	liveSet.Revision = data.Revision

	m.AddLiveSet()

	return pipeline.NewDocumentToIndexMsg(tagger.IdHash(path), liveSet)
}
