package abletonsrc

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
	"strconv"
	"strings"
)

var versionNumberRegex = regexp.MustCompile(`^(\d+\.)?(\d+\.)?(\d+)`)

func ParseLiveSet(
	stat *stats.Statistics,
	tc *tagger.TagCollector,
	path string,
	data *XmlRoot,
) *workload.DocumentPayload {
	// We only support v9+
	if !data.IsFromMinorVersion(9) {
		return nil
	}

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

	tagLiveSetFile(path, tb, doc)
	tagLiveSetTracks(data, tb, doc)
	tagLiveSetVersion(data, tb, doc)
	tagLiveSetTempo(data, tb, doc)

	doc.EngraveTags(tb)

	stat.IncrementCounter(AbletonLiveSet)

	return workload.NewDocumentPayload(doc.GetAutoId(), doc)
}

func tagLiveSetFile(path string, tb *tagger.TagBucket, doc *LiveSetDocument) {
	// Extract some details about the file itself
	fstat, err := times.Stat(path)
	if err == nil {
		// Handle the basic modification time
		year, month, _ := fstat.ModTime().Date()

		// Simple scalars
		tb.Add(fmt.Sprintf("file:mtime-year=%d", year))
		doc.Enrich(strconv.Itoa(year))

		tb.Add(fmt.Sprintf("file:mtime-weekday=%d", fstat.ModTime().Weekday()))
		doc.Enrich(fstat.ModTime().Weekday().String())

		// Month based breakdowns
		tb.Add(fmt.Sprintf("file:mtime-month=%d", month))
		doc.Enrich(month.String())

		mq := (month + 2) / 3
		tb.Add(fmt.Sprintf("file:mtime-quarter=%d", mq))
		doc.Enrich(fmt.Sprintf("Q%d", mq))

		// Week number is a bit more complex, a week can span years, but for now we just want the week number.
		_, wno := fstat.ModTime().ISOWeek()
		tb.Add(fmt.Sprintf("file:mtime-weekno=%d", wno))
		doc.Enrich(fmt.Sprintf("Week %d", wno))

		// Do the same for the creation time, if possible
		if fstat.HasBirthTime() {
			year, month, _ := fstat.BirthTime().Date()

			// Simple scalars
			tb.Add(fmt.Sprintf("file:btime-year=%d", year))
			//doc.Enrich(strconv.Itoa(year))

			tb.Add(fmt.Sprintf("file:btime-weekday=%d", fstat.BirthTime().Weekday()))
			//doc.Enrich(fstat.ModTime().Weekday().String())

			// Month based breakdowns
			tb.Add(fmt.Sprintf("file:btime-month=%d", month))
			//doc.Enrich(month.String())

			bq := (month + 2) / 3
			tb.Add(fmt.Sprintf("file:btime-quarter=%d", bq))
			//doc.Enrich(fmt.Sprintf("Q%d", bq))

			// Week number is a bit more complex, a week can span years, but for now we just want the week number.
			_, wno := fstat.BirthTime().ISOWeek()
			tb.Add(fmt.Sprintf("file:btime-weekno=%d", wno))
			//doc.Enrich(fmt.Sprintf("Week %d", wno))

			// Lets add some zodiac signs
			for _, zodiac := range gozodiacs.GetWesternZodiacsForDate(fstat.BirthTime()) {
				tb.Add(fmt.Sprintf("file:zodiac-western=%s", strings.ToLower(zodiac.String())))
				doc.Enrich(zodiac.String())
			}

			cz := gozodiacs.GetChineseZodiacSign(fstat.BirthTime()).String()
			tb.Add(fmt.Sprintf("file:zodiac-chinese=%s", strings.ToLower(cz)))
			doc.Enrich(cz)
		}
	}
}

func tagLiveSetTempo(data *XmlRoot, tb *tagger.TagBucket, doc *LiveSetDocument) {
	// @todo how to handle multi tempo files, i.e. through tempo automation?
	if data.LiveSet.MasterTrack.DeviceChain.Mixer.Tempo.Manual.Value > 0 {
		// If we have a rounded tempo, we just need to add one tag
		if math.Trunc(data.LiveSet.MasterTrack.DeviceChain.Mixer.Tempo.Manual.Value) == data.LiveSet.MasterTrack.DeviceChain.Mixer.Tempo.Manual.Value {
			bv := int(math.Round(data.LiveSet.MasterTrack.DeviceChain.Mixer.Tempo.Manual.Value))
			tb.Add(fmt.Sprintf("bpm=%d", bv))
			doc.Enrich(strconv.Itoa(bv))
		} else {
			// Otherwise it's a weird file where the tempo is a fraction, like in some XmlRoot delivered ALS files.
			// We just add both rounded values to the tb
			lv := int(math.Floor(data.LiveSet.MasterTrack.DeviceChain.Mixer.Tempo.Manual.Value))
			tb.Add(fmt.Sprintf("bpm=%d", lv))
			doc.Enrich(strconv.Itoa(lv))

			hv := int(math.Ceil(data.LiveSet.MasterTrack.DeviceChain.Mixer.Tempo.Manual.Value))
			tb.Add(fmt.Sprintf("bpm=%d", hv))
			doc.Enrich(strconv.Itoa(hv))
		}
	}
}

func tagLiveSetTracks(data *XmlRoot, tb *tagger.TagBucket, doc *LiveSetDocument) {
	// Overall track specifics
	tb.Add(fmt.Sprintf("ableton-live-set:tracks:count=%d", len(data.LiveSet.Tracks.AudioTracks)+len(data.LiveSet.Tracks.MidiTracks)))

	// Audio track specifics
	tb.Add(fmt.Sprintf("ableton-live-set:audio-tracks:count=%d", len(data.LiveSet.Tracks.AudioTracks)))
	if len(data.LiveSet.Tracks.AudioTracks) > 0 {
		tb.Add("ableton-live-set:audio-tracks:available=true")
		doc.Enrich("Audio track")
	} else {
		tb.Add("ableton-live-set:audio-tracks:available=false")
	}

	// Midi track specifics
	tb.Add(fmt.Sprintf("ableton-live-set:midi-tracks:count=%d", len(data.LiveSet.Tracks.MidiTracks)))
	if len(data.LiveSet.Tracks.MidiTracks) > 0 {
		tb.Add("ableton-live-set:midi-tracks:available=true")
		doc.Enrich("MIDI track")
	} else {
		tb.Add("ableton-live-set:midi-tracks:available=false")
		doc.Enrich("WTF")
	}
}

func tagLiveSetVersion(data *XmlRoot, tb *tagger.TagBucket, doc *LiveSetDocument) {
	// Extract software version
	if strings.HasPrefix(data.Creator, "Ableton Live ") {
		rawVersion := strings.TrimPrefix(data.Creator, "Ableton Live ")

		tb.Add(fmt.Sprintf("ableton:version=%s", rawVersion))

		if versionNumberRegex.MatchString(rawVersion) {
			verParts := strings.Split(versionNumberRegex.FindString(rawVersion), ".")

			// Major version tag
			majorV := strings.Join(verParts[:1], ".")
			tb.Add(fmt.Sprintf("ableton:version=%s", majorV))
			doc.Enrich(fmt.Sprintf("v" + majorV))

			// Minor version tag
			minorV := strings.Join(verParts[:2], ".")
			tb.Add(fmt.Sprintf("ableton:version=%s", minorV))

			// Patch version tag, just to be sure, so that "11.1.5d1" also shows up.
			if len(verParts) == 3 {
				patchV := strings.Join(verParts[:3], ".")
				tb.Add(fmt.Sprintf("ableton:version=%s", patchV))
			}
		}
	}
}
