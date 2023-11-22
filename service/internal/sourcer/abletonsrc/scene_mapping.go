package abletonsrc

import (
	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/mapping"
)

func buildSceneMapping() *mapping.DocumentMapping {
	m := bleve.NewDocumentMapping()

	createBaseMappings(m)
	createUserNameMapping(m)
	createColorMapping(m)
	createUserInfoTextMapping(m)
	createTempoWithToggle(m)

	return m
}
