package abletonv5

import (
	"github.com/adrianrudnik/ablegram/internal/pipeline"
	"github.com/adrianrudnik/ablegram/internal/stats"
	"github.com/adrianrudnik/ablegram/internal/tagger"
)

func ParseClips(stat *stats.Statistics, path string, data *XmlRoot) []*pipeline.DocumentToIndexMsg {
	docs := make([]*pipeline.DocumentToIndexMsg, 0, 10)

	for _, chain := range data.LiveSet.GetAllTrackDeviceChains() {
		for _, slots := range chain.MainSequencer.ClipSlotList.ClipSlots {
			if slots.MidiClip != nil {
				tags := tagger.NewTagger()
				tags.Add("type:ableton-midi-clip")

				doc := NewMidiClipDocument()
				doc.LoadDisplayName([]string{
					slots.MidiClip.Name.Value,
				})
				doc.LoadUserName(slots.MidiClip.Name.Value, tags)
				doc.LoadUserInfoText(slots.MidiClip.Annotation.Value, tags)
				doc.LoadFileReference(path, tags)
				doc.LoadColor(slots.MidiClip.Color.Value, tags)
				// @todo LoadScaleInformation

				doc.EngraveTags(tags)

				docs = append(docs, pipeline.NewDocumentToIndexMsg(doc.GetAutoId(), doc))

				stat.IncrementCounter(AbletonMidiClip)
			}

			if slots.AudioClip != nil {
				tags := tagger.NewTagger()
				tags.Add("type:ableton-midi-clip")

				doc := NewAudioClipDocument()
				doc.LoadDisplayName([]string{
					slots.AudioClip.Name.Value,
				})
				doc.LoadUserName(slots.AudioClip.Name.Value, tags)
				doc.LoadUserInfoText(slots.AudioClip.Annotation.Value, tags)
				doc.LoadFileReference(path, tags)

				doc.EngraveTags(tags)

				docs = append(docs, pipeline.NewDocumentToIndexMsg(doc.GetAutoId(), doc))

				stat.IncrementCounter(AbletonAudioClip)
			}
		}
	}

	return docs
}
