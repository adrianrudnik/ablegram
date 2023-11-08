package tagger

import (
	"github.com/adrianrudnik/ablegram/internal/config"
	"github.com/adrianrudnik/ablegram/internal/pusher"
	"github.com/adrianrudnik/ablegram/internal/workload"
	"github.com/samber/lo"
	"strings"
	"sync"
	"time"
)

type TagCollector struct {
	config *config.Config

	baseTags  map[string]uint64
	baseMutex sync.RWMutex

	detailedTags  map[string]uint64
	detailedMutex sync.RWMutex

	TriggerUpdate func()
}

func NewTagCollector(conf *config.Config) *TagCollector {
	b := &TagCollector{
		config:        conf,
		baseTags:      make(map[string]uint64, 500),
		detailedTags:  make(map[string]uint64, 2000),
		TriggerUpdate: func() {},
	}

	return b
}

func (c *TagCollector) WirePusher(pushChan chan<- workload.PushMessage) {
	c.TriggerUpdate, _ = lo.NewDebounce(250*time.Millisecond, func() {
		pushChan <- pusher.NewTagUpdatePush(c.GetDetailedTags())
	})
}

func (c *TagCollector) NewBucket() *TagBucket {
	return NewTagBucket(c)
}

func (c *TagCollector) collectBaseTag(t string) {
	c.baseMutex.Lock()
	defer c.baseMutex.Unlock()

	// Extract a variant of the tag without a value and increment the baseTags counter
	if strings.Contains(t, "=") {
		parts := strings.Split(t, "=")

		if _, ok := c.baseTags[parts[0]]; ok {
			t = parts[0]
		} else {
			t = parts[0]
		}
	}

	if _, ok := c.baseTags[t]; ok {
		c.baseTags[t]++
	} else {
		c.baseTags[t] = 1
	}

	c.TriggerUpdate()
}

func (c *TagCollector) collectDetailedTag(t string) {
	c.detailedMutex.Lock()
	defer c.detailedMutex.Unlock()

	if _, ok := c.detailedTags[t]; ok {
		c.detailedTags[t]++
	} else {
		c.detailedTags[t] = 1
	}

	c.TriggerUpdate()
}

func (c *TagCollector) GetBaseTags() map[string]uint64 {
	c.baseMutex.RLock()
	defer c.baseMutex.RUnlock()

	// Create a copy and return that
	tags := make(map[string]uint64, len(c.baseTags))
	for k, v := range c.baseTags {
		tags[k] = v
	}

	return tags
}

func (c *TagCollector) GetDetailedTags() map[string]uint64 {
	c.detailedMutex.RLock()
	defer c.detailedMutex.RUnlock()

	// Create a copy and return that
	tags := make(map[string]uint64, len(c.detailedTags))
	for k, v := range c.detailedTags {
		tags[k] = v
	}

	return tags
}
