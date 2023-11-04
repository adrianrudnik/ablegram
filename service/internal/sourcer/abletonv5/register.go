package abletonv5

import "github.com/blevesearch/bleve/v2/mapping"

func RegisterToIndex(idx *mapping.IndexMappingImpl) {
	idx.AddDocumentMapping("LiveSet", buildLiveSetMapping())
	idx.AddDocumentMapping("AudioTrack", buildAudioTrackMapping())
	idx.AddDocumentMapping("MidiTrack", buildMidiTrackMapping())
	idx.AddDocumentMapping("ReturnTrack", buildReturnTrackMapping())
	idx.AddDocumentMapping("Clip", buildClipMapping())
}
