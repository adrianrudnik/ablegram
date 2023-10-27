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
	"strings"
)

func ParseLiveSet(m *stats.Metrics, path string, data *ablv5schema.Ableton) *pipeline.DocumentToIndexMsg {
	tags := tagger.NewTagger()

	if util.PathContainsFolder(path, "Live Recordings") {
		tags.AddSystemTag("location:live-recording")
	}

	if util.PathContainsFolder(path, "Trash") || util.PathContainsFolder(path, "$Recycle.Bin") {
		tags.AddSystemTag("location:trash")
	} else if util.PathContainsFolder(path, "Factory Packs") {
		tags.AddSystemTag("location:factory-pack")
	} else if util.PathContainsFolder(path, "Cloud Manager") {
		tags.AddSystemTag("location:cloud-manager")
	} else if util.PathContainsFolder(path, "User Library") {
		tags.AddSystemTag("location:user-library")
	} else {
		tags.AddSystemTag("location:elsewhere")
	}

	// @todo Factory Preset, User Preset, User Library, Factory Library
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

	if len(data.LiveSet.Tracks.AudioTracks) > 0 && len(data.LiveSet.Tracks.MidiTracks) > 0 {
		tags.AddSystemTag("live-set:has-midi-audio-tracks")
	}

	if strings.HasPrefix(data.Creator, "Ableton Live ") {
		tags.AddSystemTag(fmt.Sprintf("ableton:version:%s", strings.TrimPrefix(data.Creator, "Ableton Live ")))
	}

	if data.LiveSet.MasterTrack.DeviceChain.Mixer.Tempo.Manual.Value > 0 {
		tags.AddSystemTag(fmt.Sprintf("live-set:tempo:%d", data.LiveSet.MasterTrack.DeviceChain.Mixer.Tempo.Manual.Value))
	}

	liveSet := indexer.NewLiveSetDocument()
	liveSet.Tags = tags.GetAllAndClear()
	liveSet.DisplayName = filepath.Base(path)
	liveSet.Filename = path
	liveSet.MajorVersion = data.MajorVersion
	liveSet.MinorVersion = data.MinorVersion
	liveSet.Creator = data.Creator
	liveSet.Revision = data.Revision
	liveSet.ScaleRoot = data.LiveSet.ScaleInformation.HumanizeRootNote()
	liveSet.ScaleName = data.LiveSet.ScaleInformation.Name.Value
	liveSet.Scale = fmt.Sprintf("%s %s", liveSet.ScaleRoot, liveSet.ScaleName)
	liveSet.InKey = data.LiveSet.InKey.Value
	liveSet.Tempo = data.LiveSet.MasterTrack.DeviceChain.Mixer.Tempo.Manual.Value

	m.AddLiveSet()

	return pipeline.NewDocumentToIndexMsg(tagger.IdHash(path), liveSet)
}
