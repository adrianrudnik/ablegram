package abletonsrc

import (
	"encoding/xml"
	"fmt"
	"github.com/adrianrudnik/ablegram/internal/stats"
	"github.com/adrianrudnik/ablegram/internal/tagger"
	"github.com/adrianrudnik/ablegram/internal/workload"
	"github.com/antchfx/xmlquery"
	"github.com/samber/lo"
	"strings"
)

func ParseInfotext(
	stat *stats.Statistics,
	tc *tagger.TagCollector,
	path string,
	data *XmlRoot,
) []*workload.DocumentPayload {
	docs := make([]*workload.DocumentPayload, 0, 100)

	// Works with v9+
	if !data.IsFromMinorVersion(9) {
		return docs
	}

	// We use a xpath based query to find all occurrences of the infotext entries
	xqDoc, err := xmlquery.Parse(strings.NewReader(data.Raw))
	if err != nil {
		return docs
	}

	for _, n := range xmlquery.Find(xqDoc, "//Annotation") {
		var it XmlStringValue
		err := xml.Unmarshal([]byte(n.OutputXML(true)), &it)
		if err != nil {
			Logger.Warn().Err(err).Msg("Failed to parse by xpath")
			continue
		}

		// Skip empty ones
		if strings.TrimSpace(it.Value) == "" {
			continue
		}

		tb := tc.NewBucket()
		tb.Add("type:ableton-info-text")

		doc := NewInfoTextDocument()
		doc.LoadFileReference(path, tb)
		doc.LoadDisplayName([]string{})
		doc.LoadUserInfoText(it.Value, tb)

		// We recurse the tree where we found this annotation.
		// We need to clean up the result, Ableton files leaves us with things that might
		// be duplicates (like DeviceChain) so we do a general clean
		pp := strings.Split(strings.TrimLeft(extractInfotextTree(n), "/"), "/")
		doc.Parent = strings.Join(lo.Uniq(pp), "/")

		doc.EngraveTags(tb)

		docs = append(docs, workload.NewDocumentPayload(doc.GetAutoId(), doc))

		stat.IncrementCounter(AbletonInfoText)
	}

	return docs
}

func extractInfotextTree(node *xmlquery.Node) string {
	// A > B > C > D
	if node.Parent == nil {
		return node.Data
	}

	return fmt.Sprintf("%s/%s", extractInfotextTree(node.Parent), node.Data)
}
