package abletonv5

import (
	"fmt"
	"github.com/adrianrudnik/ablegram/internal/stats"
	"github.com/adrianrudnik/ablegram/internal/tagger"
	"github.com/adrianrudnik/ablegram/internal/workload"
)

func ParseScenes(
	stat *stats.Statistics,
	tc *tagger.TagCollector,
	path string,
	data *XmlRoot,
) []*workload.DocumentPayload {
	docs := make([]*workload.DocumentPayload, 0, 10)

	for _, scene := range data.LiveSet.Scenes {
		tb := tc.NewBucket()
		tb.Add("type:ableton-scene")

		doc := NewSceneDocument()
		doc.LoadDisplayName([]string{
			fmt.Sprintf("%d", scene.Id),
			scene.Name.Value,
		})
		doc.LoadUserInfoText(scene.Annotation.Value, tb)
		doc.LoadTempoWithToggle(&scene.XmlTempoWithToggleNode, tb)
		doc.LoadFileReference(path, tb)
		doc.LoadColor(scene.Color.Value, tb)

		doc.EngraveTags(tb)

		docs = append(docs, workload.NewDocumentPayload(doc.GetAutoId(), doc))

		stat.IncrementCounter(AbletonScene)
	}

	return docs
}
