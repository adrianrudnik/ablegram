package tagger

import (
	"fmt"
	"github.com/samber/lo"
	"slices"
	"strings"
)

type Tagger struct {
	tags []string
}

func NewTagger() *Tagger {
	return &Tagger{
		tags: make([]string, 0, 20),
	}
}

func (t *Tagger) AddSystemTag(tag string) {
	tag = strings.TrimSpace(tag)
	if tag == "" {
		return
	}

	t.tags = append(t.tags, fmt.Sprintf("sys:%s", tag))
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
