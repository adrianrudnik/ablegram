package abletonv5

import (
	"fmt"
	"github.com/adrianrudnik/ablegram/internal/pipeline"
	"github.com/adrianrudnik/ablegram/internal/stats"
	"github.com/adrianrudnik/ablegram/internal/tagger"
)

func ParseScenes(stat *stats.Statistics, path string, data *XmlRoot) []*pipeline.DocumentToIndexMsg {
	docs := make([]*pipeline.DocumentToIndexMsg, 0, 10)

	for _, scene := range data.LiveSet.Scenes {
		tags := tagger.NewTagger()
		tags.Add("type:ableton-scene")

		doc := NewSceneDocument()
		doc.LoadDisplayName([]string{
			fmt.Sprintf("%d", scene.Id),
			scene.Name.Value,
		})
		doc.LoadUserInfoText(scene.Annotation.Value, tags)
		doc.LoadTempoWithToggle(&scene.XmlTempoWithToggle, tags)
		doc.LoadFileReference(path, tags)
		doc.LoadColor(scene.Color.Value, tags)

		doc.EngraveTags(tags)

		docs = append(docs, pipeline.NewDocumentToIndexMsg(doc.GetAutoId(), doc))

		stat.IncrementCounter(AbletonScene)
	}

	return docs
}
