package abletonv5

import (
	"fmt"
	"github.com/adrianrudnik/ablegram/internal/pipeline"
	"github.com/adrianrudnik/ablegram/internal/stats"
	"github.com/adrianrudnik/ablegram/internal/tagger"
	"github.com/adrianrudnik/ablegram/internal/util"
	"github.com/djherbis/times"
	"github.com/duaneking/gozodiacs"
	"math"
	"path/filepath"
	"regexp"
	"strings"
)

var versionNumberRegex = regexp.MustCompile(`^(\d+\.)?(\d+\.)?(\d+)`)

func ParseLiveSet(stat *stats.Statistics, path string, data *XmlRoot) *pipeline.DocumentToIndexMsg {
	// Extract the tags for live sets
	tags := tagger.NewTagger()
	tags.Add("type:ableton-live-set")

	doc := NewLiveSetDocument()
	doc.LoadDisplayName([]string{filepath.Base(path)})
	doc.LoadFileReference(path, tags)
	doc.LoadUserInfoText(data.LiveSet.Annotation.Value, tags)

	doc.MajorVersion = data.MajorVersion
	doc.MinorVersion = data.MinorVersion
	doc.Creator = data.Creator
	doc.Revision = data.Revision

	//doc.ScaleRootNote = data.LiveSet.ScaleInformation.HumanizeRootNote()
	//doc.ScaleName = data.LiveSet.ScaleInformation.Name.Value
	//doc.Scale = fmt.Sprintf("%s %s", doc.ScaleRootNote, doc.ScaleName)

	doc.InKey = data.LiveSet.InKey.Value
	doc.Tempo = int64(math.Round(data.LiveSet.MasterTrack.DeviceChain.Mixer.Tempo.Manual.Value))

	doc.MidiTrackCount = len(data.LiveSet.Tracks.MidiTracks)
	doc.AudioTrackCount = len(data.LiveSet.Tracks.AudioTracks)

	tagLiveSetPath(path, tags)
	tagLiveSetFile(path, tags)
	tagLiveSetTracks(data, tags)
	tagLiveSetVersion(data, tags)
	tagLiveSetTempo(data, tags)

	doc.EngraveTags(tags)

	stat.IncrementCounter(AbletonLiveSet)

	return pipeline.NewDocumentToIndexMsg(doc.GetAutoId(), doc)
}

func tagLiveSetFile(path string, tags *tagger.Tagger) {
	// Extract some details about the file itself
	fstat, err := times.Stat(path)
	if err == nil {
		// Handle the basic modification time
		year, month, _ := fstat.ModTime().Date()

		// Simple scalars
		tags.Add(fmt.Sprintf("file:mtime-year=%d", year))
		tags.Add(fmt.Sprintf("file:mtime-weekday=%d", fstat.ModTime().Weekday()))

		// Month based breakdowns
		tags.Add(fmt.Sprintf("file:mtime-month=%d", month))
		tags.Add(fmt.Sprintf("file:mtime-quarter=%d", (month+2)/3))

		// Week number is a bit more complex, a week can span years, but for now we just want the week number.
		_, wno := fstat.ModTime().ISOWeek()
		tags.Add(fmt.Sprintf("file:mtime-weekno=%d", wno))

		// Do the same for the creation time, if possible
		if fstat.HasBirthTime() {
			year, month, _ := fstat.BirthTime().Date()

			// Simple scalars
			tags.Add(fmt.Sprintf("file:btime-year=%d", year))
			tags.Add(fmt.Sprintf("file:btime-weekday=%d", fstat.ModTime().Weekday()))

			// Month based breakdowns
			tags.Add(fmt.Sprintf("file:btime-month=%d", month))
			tags.Add(fmt.Sprintf("file:btime-quarter=%d", (month+2)/3))

			// Week number is a bit more complex, a week can span years, but for now we just want the week number.
			_, wno := fstat.ModTime().ISOWeek()
			tags.Add(fmt.Sprintf("file:btime-weekno=%d", wno))

			// Lets add some zodiac signs
			for _, zodiac := range gozodiacs.GetWesternZodiacsForDate(fstat.BirthTime()) {
				tags.Add(fmt.Sprintf("file:zodiac-western=%s", strings.ToLower(zodiac.String())))
			}

			tags.Add(fmt.Sprintf("file:zodiac-chinese=%s", strings.ToLower(gozodiacs.GetChineseZodiacSign(fstat.BirthTime()).String())))
		}
	}
}

func tagLiveSetTempo(data *XmlRoot, tags *tagger.Tagger) {
	// @todo how to handle multi tempo files, i.e. through tempo automation?
	if data.LiveSet.MasterTrack.DeviceChain.Mixer.Tempo.Manual.Value > 0 {
		// If we have a rounded tempo, we just need to add one tag
		if math.Trunc(data.LiveSet.MasterTrack.DeviceChain.Mixer.Tempo.Manual.Value) == data.LiveSet.MasterTrack.DeviceChain.Mixer.Tempo.Manual.Value {
			tags.Add(fmt.Sprintf("bpm=%d", int(math.Round(data.LiveSet.MasterTrack.DeviceChain.Mixer.Tempo.Manual.Value))))
		} else {
			// Otherwise it's a weird file where the tempo is a fraction, like in some XmlRoot delivered ALS files.
			// We just add both rounded values to the tags
			tags.Add(fmt.Sprintf("bpm=%d", int(math.Floor(data.LiveSet.MasterTrack.DeviceChain.Mixer.Tempo.Manual.Value))))
			tags.Add(fmt.Sprintf("bpm=%d", int(math.Ceil(data.LiveSet.MasterTrack.DeviceChain.Mixer.Tempo.Manual.Value))))
		}
	}
}

func tagLiveSetTracks(data *XmlRoot, tags *tagger.Tagger) {
	// Overall track specifics
	tags.Add(fmt.Sprintf("ableton-live-set:tracks:count=%d", len(data.LiveSet.Tracks.AudioTracks)+len(data.LiveSet.Tracks.MidiTracks)))

	// Audio track specifics
	tags.Add(fmt.Sprintf("ableton-live-set:audio-tracks:count=%d", len(data.LiveSet.Tracks.AudioTracks)))
	if len(data.LiveSet.Tracks.AudioTracks) > 0 {
		tags.Add("ableton-live-set:audio-tracks:present=true")
	} else {
		tags.Add("ableton-live-set:audio-tracks:present=false")
	}

	// Midi track specifics
	tags.Add(fmt.Sprintf("ableton-live-set:midi-tracks:count=%d", len(data.LiveSet.Tracks.MidiTracks)))
	if len(data.LiveSet.Tracks.MidiTracks) > 0 {
		tags.Add("ableton-live-set:midi-tracks:present=true")
	} else {
		tags.Add("ableton-live-set:midi-tracks:present=false")
	}
}

func tagLiveSetPath(path string, tags *tagger.Tagger) {
	simplePath := strings.ToLower(filepath.ToSlash(path))

	if util.PathContainsFolder(path, "Live Recordings") {
		tags.Add("file:location=ableton-live-recording")
	}

	if util.PathContainsFolder(path, "Trash") || util.PathContainsFolder(path, "$Recycle.Bin") {
		tags.Add("file:location=trash")
	} else if util.PathContainsFolder(path, "Factory Packs") {
		tags.Add("file:location=factory-pack")
	} else if util.PathContainsFolder(path, "Cloud Manager") {
		tags.Add("file:location=cloud-manager")
	} else if util.PathContainsFolder(path, "User Library") {
		tags.Add("file:location=user-library")
	} else if strings.Contains(simplePath, "/dropbox") {
		tags.Add("file:location=dropbox")
	} else if strings.Contains(simplePath, "/onedrive") {
		tags.Add("file:location=onedrive")
	} else if strings.Contains(simplePath, "/google drive") {
		tags.Add("file:location=google-drive")
	} else if strings.Contains(simplePath, "/pCloudDrive") {
		tags.Add("file:location=pcloud")
	} else {
		tags.Add("file:location=elsewhere")
	}
}

func tagLiveSetVersion(data *XmlRoot, tags *tagger.Tagger) {
	// Extract software version
	if strings.HasPrefix(data.Creator, "Ableton Live ") {
		rawVersion := strings.TrimPrefix(data.Creator, "Ableton Live ")

		tags.Add(fmt.Sprintf("ableton:version=%s", rawVersion))

		if versionNumberRegex.MatchString(rawVersion) {
			verParts := strings.Split(versionNumberRegex.FindString(rawVersion), ".")

			// Major version tag
			tags.Add(fmt.Sprintf("ableton:version=%s", strings.Join(verParts[:1], ".")))

			// Minor version tag
			tags.Add(fmt.Sprintf("ableton:version=%s", strings.Join(verParts[:2], ".")))

			// Patch version tag, just to be sure, so that "11.1.5d1" also shows up.
			if len(verParts) == 3 {
				tags.Add(fmt.Sprintf("ableton:version=%s", strings.Join(verParts[:3], ".")))
			}
		}
	}
}
