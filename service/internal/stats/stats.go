package stats

import (
	"github.com/adrianrudnik/ablegram/internal/pusher"
	"sync/atomic"
)

type Metrics struct {
	validFiles     atomic.Uint64
	invalidFiles   atomic.Uint64
	liveSets       atomic.Uint64
	indexDocuments atomic.Uint64
	midiTracks     atomic.Uint64
	audioTracks    atomic.Uint64
	broadcastChan  chan<- interface{}
}

func NewMetrics(broadcastChan chan<- interface{}) *Metrics {
	return &Metrics{
		broadcastChan: broadcastChan,
	}
}

func (s *Metrics) AddValidFile() {
	v := s.validFiles.Add(1)
	s.broadcastChan <- pusher.NewMetricUpdatePush("files_valid", v)
}

func (s *Metrics) AddInvalidFile() {
	v := s.invalidFiles.Add(1)
	s.broadcastChan <- pusher.NewMetricUpdatePush("files_invalid", v)
}

func (s *Metrics) AddLiveSet() {
	v := s.liveSets.Add(1)
	s.broadcastChan <- pusher.NewMetricUpdatePush("live_sets", v)
}

func (s *Metrics) SetIndexDocuments(count uint64) {
	s.indexDocuments.Store(count)
	s.broadcastChan <- pusher.NewMetricUpdatePush("index_docs", count)
}

func (s *Metrics) AddMidiTrack() {
	v := s.midiTracks.Add(1)
	s.broadcastChan <- pusher.NewMetricUpdatePush("midi_tracks", v)
}

func (s *Metrics) AddAudioTrack() {
	v := s.audioTracks.Add(1)
	s.broadcastChan <- pusher.NewMetricUpdatePush("audio_tracks", v)
}
