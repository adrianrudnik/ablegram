package parser

import (
	"bytes"
	"compress/gzip"
	"io"
	"os"
)

func extractGzip(path string) ([]byte, error) {
	// Extract the underlying XML content
	sourceFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer sourceFile.Close()

	sourceReader, err := gzip.NewReader(sourceFile)
	if err != nil {
		return nil, err
	}
	defer sourceReader.Close()

	var sourceContent bytes.Buffer
	_, err = io.Copy(&sourceContent, sourceReader)
	if err != nil {
		return nil, err
	}

	return sourceContent.Bytes(), nil
}
