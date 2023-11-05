package abletonv5

import (
	"github.com/blevesearch/bleve/v2/mapping"
	"sync/atomic"
)

var idGenerator atomic.Uint64

func RegisterToIndex(idx *mapping.IndexMappingImpl) {
	idx.AddDocumentMapping("XmlLiveSet", buildLiveSetMapping())
	idx.AddDocumentMapping("XmlAudioTrack", buildAudioTrackMapping())
	idx.AddDocumentMapping("XmlMidiTrack", buildMidiTrackMapping())
	idx.AddDocumentMapping("XmlReturnTrack", buildReturnTrackMapping())
	idx.AddDocumentMapping("XmlGroupTrack", buildGroupTrackMapping())
	idx.AddDocumentMapping("Clip", buildClipMapping())
}
