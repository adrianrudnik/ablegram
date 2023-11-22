package abletonsrc

import (
	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/mapping"
)

func buildMidiChordDeviceMapping() *mapping.DocumentMapping {
	m := bleve.NewDocumentMapping()

	createBaseMappings(m)
	createUserNameMapping(m)
	createUserInfoTextMapping(m)
	createIsFoldedMapping(m)
	createIsExpandedMapping(m)

	return m
}
