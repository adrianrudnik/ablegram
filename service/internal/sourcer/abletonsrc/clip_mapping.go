package abletonsrc

import (
	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/mapping"
)

func buildMidiClipMapping() *mapping.DocumentMapping {
	m := bleve.NewDocumentMapping()

	createBaseMappings(m)
	createUserNameMapping(m)
	createUserInfoTextMapping(m)
	createColorMapping(m)
	createScaleInformationMapping(m)
	createTimeSignatureMapping(m)

	return m
}

func buildAudioClipMapping() *mapping.DocumentMapping {
	m := bleve.NewDocumentMapping()

	createBaseMappings(m)
	createUserNameMapping(m)
	createUserInfoTextMapping(m)

	return m
}
