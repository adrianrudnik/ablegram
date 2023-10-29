package ablv5parser

import (
	"fmt"
	"github.com/adrianrudnik/ablegram/internal/indexer"
	"github.com/adrianrudnik/ablegram/internal/parser/ablv5schema"
	"github.com/adrianrudnik/ablegram/internal/pipeline"
	"github.com/adrianrudnik/ablegram/internal/stats"
	"github.com/adrianrudnik/ablegram/internal/tagger"
	"github.com/adrianrudnik/ablegram/internal/util"
	"github.com/djherbis/times"
	"github.com/duaneking/gozodiacs"
	"math"
	"path/filepath"
	"strings"
)

func ParseLiveSet(m *stats.Metrics, path string, data *ablv5schema.Ableton) *pipeline.DocumentToIndexMsg {
	// Extract the tags for live sets
	tags := tagger.NewTagger()

	simplePath := strings.ToLower(filepath.ToSlash(path))

	if util.PathContainsFolder(path, "Live Recordings") {
		tags.AddSystemTag("file:location:live-recording")
	}

	if util.PathContainsFolder(path, "Trash") || util.PathContainsFolder(path, "$Recycle.Bin") {
		tags.AddSystemTag("file:location:trash")
	} else if util.PathContainsFolder(path, "Factory Packs") {
		tags.AddSystemTag("file:location:factory-pack")
	} else if util.PathContainsFolder(path, "Cloud Manager") {
		tags.AddSystemTag("file:location:cloud-manager")
	} else if util.PathContainsFolder(path, "User Library") {
		tags.AddSystemTag("file:location:user-library")
	} else if strings.Contains(simplePath, "/dropbox") {
		tags.AddSystemTag("file:location:dropbox")
	} else if strings.Contains(simplePath, "/onedrive") {
		tags.AddSystemTag("file:location:onedrive")
	} else if strings.Contains(simplePath, "/google drive") {
		tags.AddSystemTag("file:location:google-drive")
	} else {
		tags.AddSystemTag("file:location:elsewhere")
	}

	// @todo Factory Preset, User Preset, User Library, Factory Library
	if len(data.LiveSet.Tracks.AudioTracks) > 0 {
		tags.AddSystemTag("live-set:tracks:has-audio")
	} else {
		tags.AddSystemTag("live-set:tracks:no-audio")
	}

	if len(data.LiveSet.Tracks.MidiTracks) > 0 {
		tags.AddSystemTag("live-set:tracks:has-midi")
	} else {
		tags.AddSystemTag("live-set:tracks:no-midi")
	}

	if len(data.LiveSet.Tracks.AudioTracks) > 0 && len(data.LiveSet.Tracks.MidiTracks) > 0 {
		tags.AddSystemTag("live-set:tracks:has-midi-audio")
	}

	if strings.HasPrefix(data.Creator, "Ableton Live ") {
		tags.AddSystemTag(fmt.Sprintf("ableton:version:%s", strings.TrimPrefix(data.Creator, "Ableton Live ")))
	}

	if data.LiveSet.MasterTrack.DeviceChain.Mixer.Tempo.Manual.Value > 0 {
		// If we have a rounded tempo, we just need to add one tag
		if math.Trunc(data.LiveSet.MasterTrack.DeviceChain.Mixer.Tempo.Manual.Value) == data.LiveSet.MasterTrack.DeviceChain.Mixer.Tempo.Manual.Value {
			tags.AddSystemTag(fmt.Sprintf("live-set:tempo:%d", int(math.Round(data.LiveSet.MasterTrack.DeviceChain.Mixer.Tempo.Manual.Value))))
		} else {
			// Otherwise it's a weird file where the tempo is a fraction, like in some Ableton delivered ALS files.
			// We just add both rounded values to the tags
			tags.AddSystemTag(fmt.Sprintf("live-set:tempo:%d", int(math.Floor(data.LiveSet.MasterTrack.DeviceChain.Mixer.Tempo.Manual.Value))))
			tags.AddSystemTag(fmt.Sprintf("live-set:tempo:%d", int(math.Ceil(data.LiveSet.MasterTrack.DeviceChain.Mixer.Tempo.Manual.Value))))
		}
	}

	// Extract some details about the file itself

	fstat, err := times.Stat(path)
	if err == nil {
		// Handle the basic modification time
		year, month, _ := fstat.ModTime().Date()

		// Simple scalars
		tags.AddSystemTag(fmt.Sprintf("file:mtime-year:%d", year))
		tags.AddSystemTag(fmt.Sprintf("file:mtime-weekday:%d", fstat.ModTime().Weekday()))

		// Month based breakdowns
		tags.AddSystemTag(fmt.Sprintf("file:mtime-month:%d", month))
		tags.AddSystemTag(fmt.Sprintf("file:mtime-quarter:%d", (month+2)/3))

		// Week number is a bit more complex, a week can span years, but for now we just want the week number.
		_, wno := fstat.ModTime().ISOWeek()
		tags.AddSystemTag(fmt.Sprintf("file:mtime-weekno:%d", wno))

		// Do the same for the creation time, if possible
		if fstat.HasBirthTime() {
			year, month, _ := fstat.BirthTime().Date()

			// Simple scalars
			tags.AddSystemTag(fmt.Sprintf("file:btime-year:%d", year))
			tags.AddSystemTag(fmt.Sprintf("file:btime-weekday:%d", fstat.ModTime().Weekday()))

			// Month based breakdowns
			tags.AddSystemTag(fmt.Sprintf("file:btime-month:%d", month))
			tags.AddSystemTag(fmt.Sprintf("file:btime-quarter:%d", (month+2)/3))

			// Week number is a bit more complex, a week can span years, but for now we just want the week number.
			_, wno := fstat.ModTime().ISOWeek()
			tags.AddSystemTag(fmt.Sprintf("file:btime-weekno:%d", wno))

			// Lets add some zodiac signs
			for _, zodiac := range gozodiacs.GetWesternZodiacsForDate(fstat.BirthTime()) {
				tags.AddSystemTag(fmt.Sprintf("file:zodiac-western:%s", strings.ToLower(zodiac.String())))
			}

			tags.AddSystemTag(fmt.Sprintf("file:zodiac-chinese:%s", strings.ToLower(gozodiacs.GetChineseZodiacSign(fstat.BirthTime()).String())))
		}
	}

	liveSet := indexer.NewLiveSetDocument()
	liveSet.Tags = tags.GetAllAndClear()

	liveSet.PathAbsolute = path
	liveSet.PathFolder = filepath.Dir(path)
	liveSet.Filename = filepath.Base(path)

	liveSet.DisplayName = filepath.Base(path)
	liveSet.MajorVersion = data.MajorVersion
	liveSet.MinorVersion = data.MinorVersion
	liveSet.Creator = data.Creator
	liveSet.Revision = data.Revision

	liveSet.ScaleRootNote = data.LiveSet.ScaleInformation.HumanizeRootNote()
	liveSet.ScaleName = data.LiveSet.ScaleInformation.Name.Value
	liveSet.Scale = fmt.Sprintf("%s %s", liveSet.ScaleRootNote, liveSet.ScaleName)

	liveSet.InKey = data.LiveSet.InKey.Value
	liveSet.Tempo = int64(math.Round(data.LiveSet.MasterTrack.DeviceChain.Mixer.Tempo.Manual.Value))

	liveSet.MidiTrackCount = len(data.LiveSet.Tracks.MidiTracks)
	liveSet.AudioTrackCount = len(data.LiveSet.Tracks.AudioTracks)

	m.AddLiveSet()

	return pipeline.NewDocumentToIndexMsg(tagger.IdHash(path), liveSet)
}
