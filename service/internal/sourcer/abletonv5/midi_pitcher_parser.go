package abletonv5

import (
	"github.com/adrianrudnik/ablegram/internal/pipeline"
	"github.com/adrianrudnik/ablegram/internal/stats"
	"github.com/adrianrudnik/ablegram/internal/tagger"
)

func ParseMidiPitcherDevice(
	stat *stats.Statistics,
	tc *tagger.TagCollector,
	path string,
	data *XmlRoot,
) []*pipeline.DocumentToIndexMsg {
	hits := make([]XmlMidiPitcherDevice, 0, 100)

	// Find all MidiArpeggiator devices, in all known device chains
	for _, chain := range data.LiveSet.GetAllActualDeviceChains() {
		for _, device := range chain.Devices.MidiPitcher {
			hits = append(hits, device)
		}
	}

	docs := make([]*pipeline.DocumentToIndexMsg, 0, 10)

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
		doc.LoadDeviceIsExpanded(device.IsExpanded.Value, tb)
		doc.LoadDeviceIsFolded(device.IsFolded.Value, tb)

		doc.EngraveTags(tb)

		docs = append(docs, pipeline.NewDocumentToIndexMsg(doc.GetAutoId(), doc))

		stat.IncrementCounter(AbletonMidiPitcherDevice)
	}

	return docs
}
