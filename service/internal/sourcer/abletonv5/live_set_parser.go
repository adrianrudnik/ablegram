package abletonv5

import (
	"fmt"
	"github.com/adrianrudnik/ablegram/internal/stats"
	"github.com/adrianrudnik/ablegram/internal/tagger"
	"github.com/adrianrudnik/ablegram/internal/workload"
	"github.com/djherbis/times"
	"github.com/duaneking/gozodiacs"
	"math"
	"path/filepath"
	"regexp"
	"strings"
)

var versionNumberRegex = regexp.MustCompile(`^(\d+\.)?(\d+\.)?(\d+)`)

func ParseLiveSet(
	stat *stats.Statistics,
	tc *tagger.TagCollector,
	path string,
	data *XmlRoot,
) *workload.DocumentPayload {
	// Extract the tb for live sets
	tb := tc.NewBucket()
	tb.Add("type:ableton-live-set")

	doc := NewLiveSetDocument()
	doc.LoadDisplayName([]string{strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))})
	doc.LoadFileReference(path, tb)
	doc.LoadUserInfoText(data.LiveSet.Annotation.Value, tb)
	doc.LoadScaleInformation(&data.LiveSet.ScaleInformation, tb)

	doc.MajorVersion = data.MajorVersion
	doc.MinorVersion = data.MinorVersion
	doc.Creator = data.Creator

	doc.InKey = data.LiveSet.InKey.Value
	doc.Tempo = int64(math.Round(data.LiveSet.MasterTrack.DeviceChain.Mixer.Tempo.Manual.Value))

	doc.MidiTrackCount = len(data.LiveSet.Tracks.MidiTracks)
	doc.AudioTrackCount = len(data.LiveSet.Tracks.AudioTracks)

	tagLiveSetFile(path, tb)
	tagLiveSetTracks(data, tb)
	tagLiveSetVersion(data, tb)
	tagLiveSetTempo(data, tb)

	doc.EngraveTags(tb)

	stat.IncrementCounter(AbletonLiveSet)

	return workload.NewDocumentPayload(doc.GetAutoId(), doc)
}

func tagLiveSetFile(path string, tb *tagger.TagBucket) {
	// Extract some details about the file itself
	fstat, err := times.Stat(path)
	if err == nil {
		// Handle the basic modification time
		year, month, _ := fstat.ModTime().Date()

		// Simple scalars
		tb.Add(fmt.Sprintf("file:mtime-year=%d", year))
		tb.Add(fmt.Sprintf("file:mtime-weekday=%d", fstat.ModTime().Weekday()))

		// Month based breakdowns
		tb.Add(fmt.Sprintf("file:mtime-month=%d", month))
		tb.Add(fmt.Sprintf("file:mtime-quarter=%d", (month+2)/3))

		// Week number is a bit more complex, a week can span years, but for now we just want the week number.
		_, wno := fstat.ModTime().ISOWeek()
		tb.Add(fmt.Sprintf("file:mtime-weekno=%d", wno))

		// Do the same for the creation time, if possible
		if fstat.HasBirthTime() {
			year, month, _ := fstat.BirthTime().Date()

			// Simple scalars
			tb.Add(fmt.Sprintf("file:btime-year=%d", year))
			tb.Add(fmt.Sprintf("file:btime-weekday=%d", fstat.ModTime().Weekday()))

			// Month based breakdowns
			tb.Add(fmt.Sprintf("file:btime-month=%d", month))
			tb.Add(fmt.Sprintf("file:btime-quarter=%d", (month+2)/3))

			// Week number is a bit more complex, a week can span years, but for now we just want the week number.
			_, wno := fstat.ModTime().ISOWeek()
			tb.Add(fmt.Sprintf("file:btime-weekno=%d", wno))

			// Lets add some zodiac signs
			for _, zodiac := range gozodiacs.GetWesternZodiacsForDate(fstat.BirthTime()) {
				tb.Add(fmt.Sprintf("file:zodiac-western=%s", strings.ToLower(zodiac.String())))
			}

			tb.Add(fmt.Sprintf("file:zodiac-chinese=%s", strings.ToLower(gozodiacs.GetChineseZodiacSign(fstat.BirthTime()).String())))
		}
	}
}

func tagLiveSetTempo(data *XmlRoot, tb *tagger.TagBucket) {
	// @todo how to handle multi tempo files, i.e. through tempo automation?
	if data.LiveSet.MasterTrack.DeviceChain.Mixer.Tempo.Manual.Value > 0 {
		// If we have a rounded tempo, we just need to add one tag
		if math.Trunc(data.LiveSet.MasterTrack.DeviceChain.Mixer.Tempo.Manual.Value) == data.LiveSet.MasterTrack.DeviceChain.Mixer.Tempo.Manual.Value {
			tb.Add(fmt.Sprintf("bpm=%d", int(math.Round(data.LiveSet.MasterTrack.DeviceChain.Mixer.Tempo.Manual.Value))))
		} else {
			// Otherwise it's a weird file where the tempo is a fraction, like in some XmlRoot delivered ALS files.
			// We just add both rounded values to the tb
			tb.Add(fmt.Sprintf("bpm=%d", int(math.Floor(data.LiveSet.MasterTrack.DeviceChain.Mixer.Tempo.Manual.Value))))
			tb.Add(fmt.Sprintf("bpm=%d", int(math.Ceil(data.LiveSet.MasterTrack.DeviceChain.Mixer.Tempo.Manual.Value))))
		}
	}
}

func tagLiveSetTracks(data *XmlRoot, tb *tagger.TagBucket) {
	// Overall track specifics
	tb.Add(fmt.Sprintf("ableton-live-set:tracks:count=%d", len(data.LiveSet.Tracks.AudioTracks)+len(data.LiveSet.Tracks.MidiTracks)))

	// Audio track specifics
	tb.Add(fmt.Sprintf("ableton-live-set:audio-tracks:count=%d", len(data.LiveSet.Tracks.AudioTracks)))
	if len(data.LiveSet.Tracks.AudioTracks) > 0 {
		tb.Add("ableton-live-set:audio-tracks:available=true")
	} else {
		tb.Add("ableton-live-set:audio-tracks:available=false")
	}

	// Midi track specifics
	tb.Add(fmt.Sprintf("ableton-live-set:midi-tracks:count=%d", len(data.LiveSet.Tracks.MidiTracks)))
	if len(data.LiveSet.Tracks.MidiTracks) > 0 {
		tb.Add("ableton-live-set:midi-tracks:available=true")
	} else {
		tb.Add("ableton-live-set:midi-tracks:available=false")
	}
}

func tagLiveSetVersion(data *XmlRoot, tb *tagger.TagBucket) {
	// Extract software version
	if strings.HasPrefix(data.Creator, "Ableton Live ") {
		rawVersion := strings.TrimPrefix(data.Creator, "Ableton Live ")

		tb.Add(fmt.Sprintf("ableton:version=%s", rawVersion))

		if versionNumberRegex.MatchString(rawVersion) {
			verParts := strings.Split(versionNumberRegex.FindString(rawVersion), ".")

			// Major version tag
			tb.Add(fmt.Sprintf("ableton:version=%s", strings.Join(verParts[:1], ".")))

			// Minor version tag
			tb.Add(fmt.Sprintf("ableton:version=%s", strings.Join(verParts[:2], ".")))

			// Patch version tag, just to be sure, so that "11.1.5d1" also shows up.
			if len(verParts) == 3 {
				tb.Add(fmt.Sprintf("ableton:version=%s", strings.Join(verParts[:3], ".")))
			}
		}
	}
}
