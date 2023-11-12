package tagger

import (
	"fmt"
	"github.com/samber/lo"
	"slices"
	"strings"
	"sync"
)

type TagBucket struct {
	collector *TagCollector
	mutex     sync.RWMutex
	tags      []string
}

func NewTagBucket(collector *TagCollector) *TagBucket {
	return &TagBucket{
		collector: collector,
		tags:      make([]string, 0, 20),
	}
}

func (t *TagBucket) Add(tag string) {
	// Store the trimmed tag, with value
	tag = strings.TrimSpace(tag)
	if tag == "" {
		return
	}

	t.mutex.Lock()
	defer t.mutex.Unlock()

	t.tags = append(t.tags, fmt.Sprintf("%s", tag))
}

func (t *TagBucket) GetAll() []string {
	t.mutex.RLock()
	defer t.mutex.RUnlock()

	v := lo.Uniq(t.tags)
	slices.Sort(v)

	return v
}

// Engrave the tags to the global tag counters, creates a final slice and clears the current tagger.
// The given groups are respected by the collector and bundled together, for requesting them later.
// We use the groups to collect the final tags on a by-file group.
func (t *TagBucket) Engrave(groups []string) []string {
	t.mutex.RLock()
	defer t.mutex.RUnlock()

	v := t.GetAll()

	// Collect the tags to the global tag counters
	for _, tag := range v {
		t.collector.collectBaseTag(tag, groups)
		t.collector.collectDetailedTag(tag, groups)
	}

	// Empty the current one
	t.tags = make([]string, 0, 20)

	return v
}
