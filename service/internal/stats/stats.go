package stats

import (
	"github.com/adrianrudnik/ablegram/internal/pusher"
	"github.com/samber/lo"
	"sync/atomic"
	"time"
)

type Metrics struct {
	validFiles      atomic.Uint64
	invalidFiles    atomic.Uint64
	liveSets        atomic.Uint64
	indexDocuments  atomic.Uint64
	midiTracks      atomic.Uint64
	audioTracks     atomic.Uint64
	returnTracks    atomic.Uint64
	pushHistorySize atomic.Uint64

	triggerUpdate func()
}

func NewMetrics(broadcastChan chan<- interface{}) *Metrics {
	m := &Metrics{}

	m.triggerUpdate, _ = lo.NewDebounce(50*time.Millisecond, func() {
		broadcastChan <- m.collect()
	})

	return m
}

func (s *Metrics) collect() *pusher.MetricUpdatePush {
	return pusher.NewMetricUpdatePush(map[string]uint64{
		"files_valid":   s.validFiles.Load(),
		"files_invalid": s.invalidFiles.Load(),
		"live_sets":     s.liveSets.Load(),
		"index_docs":    s.indexDocuments.Load(),
		"midi_tracks":   s.midiTracks.Load(),
		"audio_tracks":  s.audioTracks.Load(),
		"return_tracks": s.returnTracks.Load(),
	})
}

func (s *Metrics) CountValidFile() {
	s.validFiles.Add(1)
	s.triggerUpdate()
}

func (s *Metrics) CountInvalidFile() {
	s.invalidFiles.Add(1)
	s.triggerUpdate()
}

func (s *Metrics) CountLiveSet() {
	s.liveSets.Add(1)
	s.triggerUpdate()
}

func (s *Metrics) SetIndexDocuments(count uint64) {
	s.indexDocuments.Store(count)
	s.triggerUpdate()
}

func (s *Metrics) CountMidiTrack() {
	s.midiTracks.Add(1)
	s.triggerUpdate()
}

func (s *Metrics) CountAudioTrack() {
	s.audioTracks.Add(1)
	s.triggerUpdate()
}

func (s *Metrics) CountReturnTrack() {
	s.returnTracks.Add(1)
	s.triggerUpdate()
}
