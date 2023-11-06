package abletonv5

import (
	"github.com/adrianrudnik/ablegram/internal/pipeline"
	"github.com/adrianrudnik/ablegram/internal/stats"
	"github.com/adrianrudnik/ablegram/internal/tagger"
)

func ParseTrackDeviceChains(stat *stats.Statistics, path string, data *XmlRoot) []*pipeline.DocumentToIndexMsg {
	docs := make([]*pipeline.DocumentToIndexMsg, 0, 10)

	for _, dc := range data.LiveSet.GetAllTrackDeviceChains() {
		tags := tagger.NewTagger()
		tags.Add("type:ableton-device-chain")

		doc := NewDeviceChainDocument()
		doc.LoadDisplayName([]string{AbletonDeviceChain})
		doc.LoadFileReference(path, tags)

		doc.DeviceCount = dc.DeviceChain.Devices.GetCount()

		doc.EngraveTags(tags)

		docs = append(docs, pipeline.NewDocumentToIndexMsg(doc.GetAutoId(), doc))

		stat.IncrementCounter(AbletonDeviceChain)
	}

	return docs
}
