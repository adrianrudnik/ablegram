package parser

import (
	"encoding/xml"
	"github.com/adrianrudnik/ablegram/parser/ablv5schema"
	"github.com/rs/zerolog"
	"os"
)

var Logger = zerolog.New(os.Stderr).With().Timestamp().Logger()

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
	r, err := parseAlsV5(path)
	if err != nil {
		return nil, err
	}

	return r, nil
}
