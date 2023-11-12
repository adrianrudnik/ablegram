package abletonv5

import (
	"github.com/blevesearch/bleve/v2/mapping"
	"sync/atomic"
)

var idGenerator atomic.Uint64

func RegisterDocumentMappings(idx *mapping.IndexMappingImpl) {
	idx.AddDocumentMapping(AbletonLiveSet, buildLiveSetMapping())
	idx.AddDocumentMapping(AbletonMidiTrack, buildMidiTrackMapping())
	idx.AddDocumentMapping(AbletonAudioTrack, buildAudioTrackMapping())
	idx.AddDocumentMapping(AbletonReturnTrack, buildReturnTrackMapping())
	idx.AddDocumentMapping(AbletonGroupTrack, buildGroupTrackMapping())
	idx.AddDocumentMapping(AbletonPreHearTrack, buildPreHearTrackMapping())
	idx.AddDocumentMapping(AbletonMidiClip, buildMidiClipMapping())
	idx.AddDocumentMapping(AbletonAudioClip, buildAudioClipMapping())
	idx.AddDocumentMapping(AbletonMixer, buildMixerMapping())
	idx.AddDocumentMapping(AbletonDeviceChain, buildDeviceChainMapping())
	idx.AddDocumentMapping(AbletonScene, buildSceneMapping())
	idx.AddDocumentMapping(AbletonSampleReference, buildSampleReferenceMapping())

	// Devices

	idx.AddDocumentMapping(AbletonMidiArpeggiatorDevice, buildMidiArpeggiatorDeviceMapping())
	idx.AddDocumentMapping(AbletonMidiChordDevice, buildMidiChordDeviceMapping())
	idx.AddDocumentMapping(AbletonMidiPitcherDevice, buildMidiPitcherDeviceMapping())
	idx.AddDocumentMapping(AbletonMidiVelocityDevice, buildMidiVelocityDeviceMapping())

	// Source file
	idx.AddDocumentMapping(AbletonAlsFile, buildAlsFileMapping())
}
