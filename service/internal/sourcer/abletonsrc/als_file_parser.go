package abletonsrc

import (
	"github.com/adrianrudnik/ablegram/internal/stats"
	"github.com/adrianrudnik/ablegram/internal/tagger"
	"github.com/adrianrudnik/ablegram/internal/workload"
	"github.com/samber/lo"
	"path/filepath"
	"slices"
)

func ParseAlsFile(
	stat *stats.Statistics,
	tc *tagger.TagCollector,
	path string,
	data *XmlRoot,
) *workload.DocumentPayload {
	// Extract the tb for live sets
	tb := tc.NewBucket()
	tb.Add("type:ableton-als-file")
	tb.Add("type:file")

	doc := NewAlsFileDocument()
	doc.LoadDisplayName([]string{filepath.Base(path)})
	doc.LoadFileReference(path, tb)

	// Engrave as normal
	doc.EngraveTags(tb)

	// But also engrave all tags found for this specific absolute path
	doc.Tags = append(lo.MapToSlice(tc.GetGroupedTags(path), func(k string, _ uint64) string {
		return k
	}))

	doc.Tags = lo.Uniq(doc.Tags)
	slices.Sort(doc.Tags)

	stat.IncrementCounter(AbletonAlsFile)

	return workload.NewDocumentPayload(doc.GetAutoId(), doc)
}
