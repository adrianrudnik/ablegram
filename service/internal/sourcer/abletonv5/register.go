package abletonv5

import "github.com/blevesearch/bleve/v2/mapping"

func RegisterToIndex(idx *mapping.IndexMappingImpl) {
	idx.AddDocumentMapping("Clip", buildClipMapping())
	idx.AddDocumentMapping("AudioTrack", buildAudioTrackMapping())
	idx.AddDocumentMapping("LiveSet", buildLiveSetMapping())
	idx.AddDocumentMapping("MidiTrack", buildMidiTrackMapping())
}
