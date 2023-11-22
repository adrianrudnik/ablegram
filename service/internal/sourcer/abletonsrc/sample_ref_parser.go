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
