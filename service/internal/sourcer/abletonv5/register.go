package abletonv5

import (
	"github.com/blevesearch/bleve/v2/mapping"
	"sync/atomic"
)

var idGenerator atomic.Uint64

func RegisterToIndex(idx *mapping.IndexMappingImpl) {
	idx.AddDocumentMapping(AbletonLiveSet, buildLiveSetMapping())
	idx.AddDocumentMapping(AbletonMidiTrack, buildMidiTrackMapping())
	idx.AddDocumentMapping(AbletonAudioTrack, buildAudioTrackMapping())
	idx.AddDocumentMapping(AbletonReturnTrack, buildReturnTrackMapping())
	idx.AddDocumentMapping(AbletonGroupTrack, buildGroupTrackMapping())
	idx.AddDocumentMapping(AbletonPreHearTrack, buildPreHearTrackMapping())
	idx.AddDocumentMapping(AbletonClip, buildClipMapping())
	idx.AddDocumentMapping(AbletonMixer, buildMixerMapping())
	idx.AddDocumentMapping(AbletonDeviceChain, buildDeviceChainMapping())
	idx.AddDocumentMapping(AbletonScene, buildSceneMapping())
}
