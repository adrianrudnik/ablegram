package abletonv5

import (
	"encoding/xml"
	"errors"
	"github.com/adrianrudnik/ablegram/internal/stats"
	"github.com/adrianrudnik/ablegram/internal/tagger"
	"github.com/adrianrudnik/ablegram/internal/workload"
	"github.com/antchfx/xmlquery"
	"os"
	"path/filepath"
	"strings"
)

func ParseSampleReferences(
	stat *stats.Statistics,
	tc *tagger.TagCollector,
	path string,
	data *XmlRoot,
) []*workload.DocumentPayload {
	docs := make([]*workload.DocumentPayload, 0, 10)

	// Skip v10 files for now, that's some workload to work through.
	if !data.IsMinorVersion(11) {
		return docs
	}

	xqDoc, err := xmlquery.Parse(strings.NewReader(data.Raw))
	if err != nil {
		return docs
	}

	// Can be very nested
	for _, n := range xmlquery.Find(xqDoc, "//SampleRef") {
		var sampleRef XmlSampleRef11
		err := xml.Unmarshal([]byte(n.OutputXML(true)), &sampleRef)
		if err != nil {
			Logger.Warn().Err(err).Msg("Failed to parse by xpath")
			continue
		}

		// Skip all paths that are of a relative type
		if sampleRef.FileReference.RelativePathType.Value != 0 {
			continue
		}

		tb := tc.NewBucket()
		tb.Add("type:ableton-sample-reference")

		doc := NewSampleReferenceDocument()
		doc.SampleAbsPath = sampleRef.FileReference.Path.Value
		doc.SampleFilename = filepath.Base(sampleRef.FileReference.Path.Value)
		doc.SampleOriginalFileSize = sampleRef.FileReference.OriginalFileSize.Value

		doc.LoadFileReference(path, tb)
		doc.LoadDisplayName([]string{
			doc.SampleFilename,
		})

		// Check if the sample is actually unavailable
		fstat, err := os.Stat(doc.SampleAbsPath)
		if errors.Is(err, os.ErrNotExist) {
			tb.Add("ableton-sample-reference:unavailable")
		} else if err != nil {
			Logger.Warn().Err(err).
				Str("path", doc.SampleAbsPath).
				Msg("Failed to stat file for sample reference")
		} else {
			if fstat.Size() != doc.SampleOriginalFileSize {
				tb.Add("ableton-sample-reference:modified")
			}
		}

		doc.EngraveTags(tb)

		docs = append(docs, workload.NewDocumentPayload(doc.GetAutoId(), doc))

		stat.IncrementCounter(AbletonSampleReference)
	}

	return docs
}
