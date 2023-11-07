package abletonv5

import (
	"github.com/adrianrudnik/ablegram/internal/pipeline"
	"github.com/adrianrudnik/ablegram/internal/stats"
	"github.com/adrianrudnik/ablegram/internal/tagger"
)

func ParseTrackDeviceChains(
	stat *stats.Statistics,
	tc *tagger.TagCollector,
	path string,
	data *XmlRoot,
) []*pipeline.DocumentToIndexMsg {
	docs := make([]*pipeline.DocumentToIndexMsg, 0, 10)

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

		docs = append(docs, pipeline.NewDocumentToIndexMsg(doc.GetAutoId(), doc))

		stat.IncrementCounter(AbletonDeviceChain)
	}

	return docs
}
