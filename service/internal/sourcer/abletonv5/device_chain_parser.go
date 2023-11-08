package abletonv5

import (
	"github.com/adrianrudnik/ablegram/internal/stats"
	"github.com/adrianrudnik/ablegram/internal/tagger"
	"github.com/adrianrudnik/ablegram/internal/workload"
)

func ParseTrackDeviceChains(
	stat *stats.Statistics,
	tc *tagger.TagCollector,
	path string,
	data *XmlRoot,
) []*workload.DocumentPayload {
	docs := make([]*workload.DocumentPayload, 0, 10)

	for _, dc := range data.LiveSet.GetAllTrackDeviceChains() {
		tb := tc.NewBucket()
		tb.Add("type:ableton-device-chain")

		// Consider a device chain empty if no devices are present
		if dc.DeviceChain.Devices.GetCount() == 0 {
			continue
		}

		doc := NewDeviceChainDocument()
		doc.LoadDisplayName([]string{AbletonDeviceChain})
		doc.LoadFileReference(path, tb)

		doc.DeviceCount = dc.DeviceChain.Devices.GetCount()

		doc.EngraveTags(tb)

		docs = append(docs, workload.NewDocumentPayload(doc.GetAutoId(), doc))

		stat.IncrementCounter(AbletonDeviceChain)
	}

	return docs
}
