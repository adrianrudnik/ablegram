package abletonv5

import (
	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/mapping"
)

func buildMixerMapping() *mapping.DocumentMapping {
	m := bleve.NewDocumentMapping()

	createFileReferenceMappings(m)
	createUserNameMapping(m)
	createUserInfoTextMapping(m)

	return m
}