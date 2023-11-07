package abletonv5

import (
	"github.com/adrianrudnik/ablegram/internal/pipeline"
	"github.com/adrianrudnik/ablegram/internal/stats"
	"github.com/adrianrudnik/ablegram/internal/tagger"
)

func ParseMidiChordDevice(stat *stats.Statistics, path string, data *XmlRoot) []*pipeline.DocumentToIndexMsg {
	hits := make([]XmlMidiChordDevice, 0, 100)

	for _, chain := range data.LiveSet.GetAllActualDeviceChains() {
		for _, device := range chain.Devices.MidiChord {
			hits = append(hits, device)
		}
	}

	docs := make([]*pipeline.DocumentToIndexMsg, 0, 10)

	for _, device := range hits {
		tags := tagger.NewTagger()
		tags.Add("type:ableton-midi-chord-device")
		tags.Add("ableton-device:midi-chord")

		doc := NewMidiChordDeviceDocument()
		doc.LoadDisplayName([]string{
			device.UserName.Value,
		})
		doc.LoadFileReference(path, tags)
		doc.LoadUserInfoText(device.Annotation.Value, tags)
		doc.LoadDeviceIsExpanded(device.IsExpanded.Value, tags)
		doc.LoadDeviceIsFolded(device.IsFolded.Value, tags)

		doc.EngraveTags(tags)

		docs = append(docs, pipeline.NewDocumentToIndexMsg(doc.GetAutoId(), doc))

		stat.IncrementCounter(AbletonMidiArpeggiatorDevice)
	}

	return docs
}
