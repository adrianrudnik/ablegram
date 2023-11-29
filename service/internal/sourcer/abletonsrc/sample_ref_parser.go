package abletonsrc

import (
	"encoding/xml"
	"errors"
	"github.com/adrianrudnik/ablegram/internal/stats"
	"github.com/adrianrudnik/ablegram/internal/tagger"
	"github.com/adrianrudnik/ablegram/internal/util"
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

	xqDoc, err := xmlquery.Parse(strings.NewReader(data.Raw))
	if err != nil {
		return docs
	}

	// Can be very nested
	for _, n := range xmlquery.Find(xqDoc, "//SampleRef") {
		tb := tc.NewBucket()
		tb.Add("type:ableton-sample-reference")

		doc := NewSampleReferenceDocument()

		if data.IsFromMinorVersion(11) {
			var sampleRef11 XmlSampleRef11

			err := xml.Unmarshal([]byte(n.OutputXML(true)), &sampleRef11)
			if err != nil {
				Logger.Warn().Err(err).Msg("Failed to parse by xpath")
				continue
			}

			// Skip all paths that are of a relative type
			/* Past me: why?
			if sampleRef11.FileReference.RelativePathType.Value != 0 {
				continue
			}
			*/

			doc.SampleAbsPath = sampleRef11.FileReference.Path.Value
			doc.SampleFilename = filepath.Base(sampleRef11.FileReference.Path.Value)
			doc.SampleOriginalFileSize = sampleRef11.FileReference.OriginalFileSize.Value

			doc.Enrich(doc.SampleFilename)
			doc.Enrich(strings.TrimSuffix(doc.SampleFilename, filepath.Ext(doc.SampleFilename)))

			doc.SampleAbsPath = sampleRef11.FileReference.Path.Value
			doc.SampleFilename = filepath.Base(sampleRef11.FileReference.Path.Value)
			doc.SampleOriginalFileSize = sampleRef11.FileReference.OriginalFileSize.Value

			doc.Enrich(doc.SampleFilename)
			doc.Enrich(strings.TrimSuffix(doc.SampleFilename, filepath.Ext(doc.SampleFilename)))
		} else {
			var sampleRef9 XmlSampleRef9

			err := xml.Unmarshal([]byte(n.OutputXML(true)), &sampleRef9)
			if err != nil {
				Logger.Warn().Err(err).Msg("Failed to parse by xpath")
				continue
			}

			// Let's prefer the SearchHint in this version, it contains more valuable information
			doc.SampleAbsPath = strings.Join(sampleRef9.FileReference.SearchHint.PathHintFolders(), "/")
			doc.SampleFilename = filepath.Base(sampleRef9.FileReference.Name.Value)
			doc.SampleOriginalFileSize = sampleRef9.FileReference.SearchHint.FileSize.Value

			doc.Enrich(doc.SampleFilename)
			doc.Enrich(strings.TrimSuffix(doc.SampleFilename, filepath.Ext(doc.SampleFilename)))

			for _, p := range sampleRef9.FileReference.RelativePathFolders() {
				doc.Enrich(p)
			}

			for _, p := range sampleRef9.FileReference.SearchHint.PathHintFolders() {
				doc.Enrich(p)
			}
		}

		doc.LoadFileReference(path, tb)
		doc.LoadDisplayName([]string{
			doc.SampleFilename,
		})

		// Check if the sample is actually unavailable
		fstat, err := os.Stat(doc.SampleAbsPath)
		if errors.Is(err, os.ErrNotExist) {
			tb.Add("ableton-sample-reference:available=false")

			// There are several scenarios that can be at play here:
			// First lets check if we might have the wrong OS
			if !util.IsPathOriginFromTheSameOs(doc.SampleAbsPath) {
				tb.Add("ableton-sample-reference:wrong-os")
			}

			// Now lets check if the files might have been collected, so it
			// has to be in the same folder as the ALS file.
			fstat, err = os.Stat(filepath.Join(filepath.Dir(path), filepath.Base(doc.SampleAbsPath)))
			if errors.Is(err, os.ErrNotExist) {
				// We can't find the file in the same folder as the ALS file
				tb.Add("ableton-sample-reference:collected=false")
			} else if err != nil {
				// Unexpected error
				Logger.Warn().Err(err).
					Str("path", doc.SampleAbsPath).
					Msg("Failed to stat collected file for sample reference")
				tb.Add("ableton-sample-reference:caution")
			} else {
				// It is available as collected, but maybe it was modified?
				tb.Add("ableton-sample-reference:collected=true")

				if fstat.Size() != doc.SampleOriginalFileSize {
					tb.Add("ableton-sample-reference:modified=true")
				} else {
					tb.Add("ableton-sample-reference:modified=false")
				}
			}
		} else if err != nil {
			Logger.Warn().Err(err).
				Str("path", doc.SampleAbsPath).
				Msg("Failed to stat file for sample reference")
			tb.Add("ableton-sample-reference:caution")
		} else {
			tb.Add("ableton-sample-reference:available=true")

			if fstat.Size() != doc.SampleOriginalFileSize {
				tb.Add("ableton-sample-reference:modified=true")
			} else {
				tb.Add("ableton-sample-reference:modified=false")
			}
		}

		doc.EngraveTags(tb)

		docs = append(docs, workload.NewDocumentPayload(doc.GetAutoId(), doc))

		stat.IncrementCounter(AbletonSampleReference)
	}

	return docs
}
