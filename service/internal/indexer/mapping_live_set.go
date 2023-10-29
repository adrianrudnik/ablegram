package indexer

import (
	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/mapping"
)

type LiveSetDocument struct {
	T    string   `json:"type"`
	Tags []string `json:"tags,omitempty"`

	DisplayName  string `json:"displayName,omitempty"`
	Filename     string `json:"filename,omitempty"`
	AbsPath      string `json:"path,omitempty"`
	MajorVersion string `json:"majorVersion,omitempty"`
	MinorVersion string `json:"minorVersion,omitempty"`
	Creator      string `json:"Creator,omitempty"`
	Revision     string `json:"Revision,omitempty"`

	Scale         string `json:"scale,omitempty"`
	ScaleName     string `json:"scaleName,omitempty"`
	ScaleRootNote string `json:"scaleRootNote,omitempty"`

	InKey bool  `json:"inKey,omitempty"`
	Tempo int64 `json:"tempo,omitempty"`

	ScaleInformation string ``
}

func (d *LiveSetDocument) Type() string {
	return d.T
}

func NewLiveSetDocument() *LiveSetDocument {
	return &LiveSetDocument{
		T: "LiveSet",
	}
}

func buildLiveSetMapping() *mapping.DocumentMapping {
	m := bleve.NewDocumentMapping()
	m.AddFieldMappingsAt("type", mapping.NewKeywordFieldMapping())
	m.AddFieldMappingsAt("tags", mapping.NewKeywordFieldMapping())

	m.AddFieldMappingsAt("displayName", NewPayloadFieldMapping())
	m.AddFieldMappingsAt("filename", NewFileFieldMapping())
	m.AddFieldMappingsAt("path", NewFulltextTextFieldMapping(false))
	m.AddFieldMappingsAt("majorVersion", NewPayloadFieldMapping())
	m.AddFieldMappingsAt("minorVersion", NewPayloadFieldMapping())
	m.AddFieldMappingsAt("creator", mapping.NewKeywordFieldMapping())
	m.AddFieldMappingsAt("revision", NewPayloadFieldMapping())

	m.AddFieldMappingsAt("scale", mapping.NewKeywordFieldMapping())
	m.AddFieldMappingsAt("scaleName", mapping.NewKeywordFieldMapping())
	m.AddFieldMappingsAt("scaleRootNote", mapping.NewKeywordFieldMapping())

	m.AddFieldMappingsAt("inKey", mapping.NewBooleanFieldMapping())
	m.AddFieldMappingsAt("tempo", mapping.NewNumericFieldMapping())

	return m
}
