package abletonv5

import (
	"github.com/adrianrudnik/ablegram/internal/pipeline"
	"github.com/adrianrudnik/ablegram/internal/stats"
	"github.com/adrianrudnik/ablegram/internal/tagger"
)

func ParseMidiVelocityDevice(stat *stats.Statistics, path string, data *XmlRoot) []*pipeline.DocumentToIndexMsg {
	hits := make([]XmlMidiVelocityDevice, 0, 100)

	// Find all MidiArpeggiator devices, in all known device chains
	for _, chain := range data.LiveSet.GetAllActualDeviceChains() {
		for _, device := range chain.Devices.MidiVelocity {
			hits = append(hits, device)
		}
	}

	docs := make([]*pipeline.DocumentToIndexMsg, 0, 10)

	for _, device := range hits {
		tags := tagger.NewTagger()
		tags.Add("type:ableton-midi-velocity-device")
		tags.Add("ableton-device:midi-velocity")

		doc := NewMidiVelocityDeviceDocument()
		doc.LoadDisplayName([]string{
			device.UserName.Value,
		})
		doc.LoadFileReference(path, tags)
		doc.LoadUserInfoText(device.Annotation.Value, tags)

		doc.LoadDeviceIsExpanded(device.IsExpanded.Value, tags)
		doc.LoadDeviceIsFolded(device.IsFolded.Value, tags)

		doc.EngraveTags(tags)

		docs = append(docs, pipeline.NewDocumentToIndexMsg(doc.GetAutoId(), doc))

		stat.IncrementCounter(AbletonMidiVelocityDevice)
	}

	return docs
}
