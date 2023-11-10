package abletonv5

import (
	"github.com/adrianrudnik/ablegram/internal/stats"
	"github.com/adrianrudnik/ablegram/internal/tagger"
	"github.com/adrianrudnik/ablegram/internal/workload"
)

func ParseMidiPitcherDevice(
	stat *stats.Statistics,
	tc *tagger.TagCollector,
	path string,
	data *XmlRoot,
) []*workload.DocumentPayload {
	hits := make([]XmlMidiPitcherDevice, 0, 100)

	// Find all MidiArpeggiator devices, in all known device chains
	for _, chain := range data.LiveSet.GetAllActualDeviceChains() {
		for _, device := range chain.Devices.MidiPitcher {
			hits = append(hits, device)
		}
	}

	docs := make([]*workload.DocumentPayload, 0, 10)

	for _, device := range hits {
		tb := tc.NewBucket()
		tb.Add("type:ableton-midi-pitcher-device")
		tb.Add("ableton-device:midi-pitcher")

		doc := NewMidiPitcherDeviceDocument()
		doc.LoadDisplayName([]string{
			device.UserName.Value,
		})
		doc.LoadFileReference(path, tb)
		doc.LoadUserInfoText(device.Annotation.Value, tb)
		doc.LoadOptionExpanded(device.IsExpanded.Value, tb)
		doc.LoadOptionFolded(device.IsFolded.Value, tb)

		doc.EngraveTags(tb)

		docs = append(docs, workload.NewDocumentPayload(doc.GetAutoId(), doc))

		stat.IncrementCounter(AbletonMidiPitcherDevice)
	}

	return docs
}
