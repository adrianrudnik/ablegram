package tagger

import (
	"fmt"
	"github.com/samber/lo"
	"slices"
	"strings"
)

type TagBucket struct {
	collector *TagCollector
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

	t.tags = append(t.tags, fmt.Sprintf("%s", tag))
}

func (t *TagBucket) GetAll() []string {
	v := lo.Uniq(t.tags)
	slices.Sort(v)

	return v
}

// Engrave the tags to the global tag counters, creates a final slice and clears the current tagger
func (t *TagBucket) Engrave() []string {
	v := t.GetAll()

	// Collect the tags to the global tag counters
	for _, tag := range v {
		t.collector.collectBaseTag(tag)
		t.collector.collectDetailedTag(tag)
	}

	// Empty the current one
	t.tags = make([]string, 0, 20)

	return v
}
