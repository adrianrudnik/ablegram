package indexer

import (
	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/mapping"
)

type LiveSetDocument struct {
	T    string   `json:"type"`
	Tags []string `json:"tags,omitempty"`

	DisplayName  string `json:"display_name,omitempty"`
	Filename     string `json:"filename,omitempty"`
	MajorVersion string `json:"major_version,omitempty"`
	MinorVersion string `json:"minor_version,omitempty"`
	Creator      string `json:"creator,omitempty"`
	Revision     string `json:"revision,omitempty"`

	ScaleRoot string `json:"scale_root_note,omitempty"`
	ScaleName string `json:"scale_name,omitempty"`
	Scale     string `json:"scale,omitempty"`

	InKey bool  `json:"in_key,omitempty"`
	Tempo int64 `json:"tempo,omitempty"`

	ScaleInformation string ``
}

func (d *LiveSetDocument) Type() string {
	return d.T
}

func NewLiveSetDocument() *LiveSetDocument {
	return &LiveSetDocument{
		T: "live_set",
	}
}

func buildLiveSetMapping(options *SearchOptions) *mapping.DocumentMapping {
	m := bleve.NewDocumentMapping()
	m.AddFieldMappingsAt("type", NewTypeFieldMapping(options))
	m.AddFieldMappingsAt("tags", NewTagFieldMapping(options))

	m.AddFieldMappingsAt("display_name", NewInfoTextFieldMapping(options))
	m.AddFieldMappingsAt("filename", NewFileFieldMapping(options))
	m.AddFieldMappingsAt("major_version", NewInfoTextFieldMapping(options))
	m.AddFieldMappingsAt("minor_version", NewInfoTextFieldMapping(options))
	m.AddFieldMappingsAt("creator", NewTagFieldMapping(options))
	m.AddFieldMappingsAt("revision", NewInfoTextFieldMapping(options))

	m.AddFieldMappingsAt("scale_root", mapping.NewKeywordFieldMapping())
	m.AddFieldMappingsAt("scale_name", mapping.NewKeywordFieldMapping())
	m.AddFieldMappingsAt("scale", mapping.NewKeywordFieldMapping())

	m.AddFieldMappingsAt("in_key", mapping.NewBooleanFieldMapping())
	m.AddFieldMappingsAt("tempo", mapping.NewNumericFieldMapping())

	return m
}
