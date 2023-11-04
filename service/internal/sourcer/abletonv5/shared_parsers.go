package abletonv5

import (
	"fmt"
	"github.com/adrianrudnik/ablegram/internal/tagger"
	"github.com/adrianrudnik/ablegram/internal/util"
)

// parseAnnotation parses the given annotation field, applies default tags and returns the annotation value.
func parseAnnotation(t *tagger.Tagger, v string) string {
	m, empty := util.EvaluateUserInput(v)

	if !empty {
		t.AddSystemTag("info:has-user-memo")
	} else {
		t.AddSystemTag("info:no-user-memo")
	}

	return m
}

// parseColor parses the given color field, applies default tags and returns the color value.
func parseColor(t *tagger.Tagger, v int16) int16 {
	t.AddSystemTag(fmt.Sprintf("color:ableton:%d", v))
	return v
}
