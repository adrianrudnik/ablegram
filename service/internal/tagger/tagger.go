package tagger

import (
	"fmt"
	"github.com/samber/lo"
	"slices"
	"strings"
	"sync"
)

var baseTags = make(map[string]uint64, 500)
var baseMutex sync.RWMutex

var detailedTags = make(map[string]uint64, 2000)
var detailedMutex sync.RWMutex

type Tagger struct {
	tags []string
}

func NewTagger() *Tagger {
	return &Tagger{
		tags: make([]string, 0, 20),
	}
}

func (t *Tagger) Add(tag string) {
	tag = strings.TrimSpace(tag)
	if tag == "" {
		return
	}

	t.tags = append(t.tags, fmt.Sprintf("%s", tag))

	collectBaseTag(tag)
	collectDetailedTag(tag)
}

func (t *Tagger) GetAll() []string {
	v := lo.Uniq(t.tags)
	slices.Sort(v)

	return v
}

func (t *Tagger) GetAllAndClear() []string {
	v := t.GetAll()
	t.tags = make([]string, 0, 20)

	return v
}

func collectBaseTag(t string) {
	baseMutex.Lock()
	defer baseMutex.Unlock()

	// Extract a variant of the tag without a value and increment the baseTags counter
	if strings.Contains(t, "=") {
		parts := strings.Split(t, "=")

		if _, ok := baseTags[parts[0]]; ok {
			t = parts[0]
		} else {
			t = parts[0]
		}
	}

	if _, ok := baseTags[t]; ok {
		baseTags[t]++
	} else {
		baseTags[t] = 1
	}
}

func collectDetailedTag(t string) {
	detailedMutex.Lock()
	defer detailedMutex.Unlock()

	if _, ok := detailedTags[t]; ok {
		detailedTags[t]++
	} else {
		detailedTags[t] = 1
	}
}

func GetBaseTags() map[string]uint64 {
	baseMutex.RLock()
	defer baseMutex.RUnlock()

	return baseTags
}

func GetDetailedTags() map[string]uint64 {
	detailedMutex.RLock()
	defer detailedMutex.RUnlock()

	return detailedTags
}
