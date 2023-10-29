package indexer

import (
	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/mapping"
)

type LiveSetDocument struct {
	T    string   `json:"type"`
	Tags []string `json:"tags,omitempty"`

	PathAbsolute string `json:"pathAbsolute,omitempty"`
	PathFolder   string `json:"pathFolder,omitempty"`
	Filename     string `json:"filename,omitempty"`

	DisplayName  string `json:"displayName,omitempty"`
	MajorVersion string `json:"majorVersion,omitempty"`
	MinorVersion string `json:"minorVersion,omitempty"`
	Creator      string `json:"creator,omitempty"`
	Revision     string `json:"revision,omitempty"`

	Scale         string `json:"scale,omitempty"`
	ScaleName     string `json:"scaleName,omitempty"`
	ScaleRootNote string `json:"scaleRootNote,omitempty"`

	InKey bool  `json:"inKey,omitempty"`
	Tempo int64 `json:"tempo,omitempty"`

	MidiTrackCount  int `json:"midiTrackCount,omitempty"`
	AudioTrackCount int `json:"audioTrackCount,omitempty"`

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

	m.AddFieldMappingsAt("pathFolder", mapping.NewKeywordFieldMapping())
	m.AddFieldMappingsAt("pathAbsolute", mapping.NewKeywordFieldMapping())
	m.AddFieldMappingsAt("filename", mapping.NewKeywordFieldMapping())

	m.AddFieldMappingsAt("displayName", NewFulltextTextFieldMapping(true))
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

	m.AddFieldMappingsAt("midiTrackCount", mapping.NewNumericFieldMapping())
	m.AddFieldMappingsAt("audioTrackCount", mapping.NewNumericFieldMapping())

	return m
}
