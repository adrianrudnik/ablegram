package abletonv5

import (
	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/mapping"
)

func buildMidiPitcherDeviceMapping() *mapping.DocumentMapping {
	m := bleve.NewDocumentMapping()

	createBaseMappings(m)
	createUserNameMapping(m)
	createUserInfoTextMapping(m)
	createIsExpandedMapping(m)
	createIsFoldedMapping(m)

	return m
}
