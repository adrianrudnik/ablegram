package abletonv5

import (
	"fmt"
	"github.com/adrianrudnik/ablegram/internal/pipeline"
	"github.com/adrianrudnik/ablegram/internal/stats"
	"github.com/adrianrudnik/ablegram/internal/tagger"
)

func ParseClips(
	stat *stats.Statistics,
	tc *tagger.TagCollector,
	path string,
	data *XmlRoot,
) []*pipeline.DocumentToIndexMsg {
	docs := make([]*pipeline.DocumentToIndexMsg, 0, 10)

	for _, chain := range data.LiveSet.GetAllTrackDeviceChains() {
		for _, slots := range chain.MainSequencer.ClipSlotList.ClipSlots {
			if slots.MidiClip != nil {
				tb := tc.NewBucket()
				tb.Add("type:ableton-midi-clip")

				doc := NewMidiClipDocument()
				doc.LoadDisplayName([]string{
					slots.MidiClip.Name.Value,
				})
				doc.LoadUserName(slots.MidiClip.Name.Value, tb)
				doc.LoadUserInfoText(slots.MidiClip.Annotation.Value, tb)
				doc.LoadFileReference(path, tb)
				doc.LoadColor(slots.MidiClip.Color.Value, tb)
				doc.LoadTimeSignature(&slots.MidiClip.TimeSignature, tb)
				doc.LoadScaleInformation(&slots.MidiClip.ScaleInformation, tb)

				// Parse midi notes
				hasProbability := false
				for _, note := range slots.MidiClip.Notes.KeyTracks {
					// Add the used notes
					tb.Add(fmt.Sprintf("note=%s", note.MidiKey.HumanReadable()))
					tb.Add(fmt.Sprintf("midi:key=%d", note.MidiKey.Value))

					for _, midiNote := range note.Notes {
						if midiNote.Probability < 1 {
							hasProbability = true
						}
					}
				}

				if hasProbability {
					tb.Add("ableton:feature:probability=true")
				}

				doc.EngraveTags(tb)

				docs = append(docs, pipeline.NewDocumentToIndexMsg(doc.GetAutoId(), doc))

				stat.IncrementCounter(AbletonMidiClip)
			}

			if slots.AudioClip != nil {
				tb := tc.NewBucket()
				tb.Add("type:ableton-midi-clip")

				doc := NewAudioClipDocument()
				doc.LoadDisplayName([]string{
					slots.AudioClip.Name.Value,
				})
				doc.LoadUserName(slots.AudioClip.Name.Value, tb)
				doc.LoadUserInfoText(slots.AudioClip.Annotation.Value, tb)
				doc.LoadFileReference(path, tb)

				doc.EngraveTags(tb)

				docs = append(docs, pipeline.NewDocumentToIndexMsg(doc.GetAutoId(), doc))

				stat.IncrementCounter(AbletonAudioClip)
			}
		}
	}

	return docs
}
