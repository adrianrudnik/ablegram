package parser

import (
	"encoding/xml"
	"git.jetbrains.space/the/ablegram/parser/ablv5schema"
)

func parseAlsV5(path string) (*Result, error) {
	rawContent, err := extractGzip(path)
	if err != nil {
		return nil, err
	}

	var data ablv5schema.Ableton

	err = xml.Unmarshal(rawContent, &data)
	if err != nil {
		return nil, err
	}

	return &Result{}, nil
}

func ParseAls(path string) (*Result, error) {
	return parseAlsV5(path)
}
