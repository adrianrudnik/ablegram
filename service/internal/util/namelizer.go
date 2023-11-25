package util

import (
	"github.com/samber/lo"
	"regexp"
	"strings"
)

var displayNameSanitizer = regexp.MustCompile(`[^a-zA-Z0-9\s]`)

func Namelize(parts []string) string {
	parts = lo.Filter(parts, func(x string, index int) bool {
		return x != ""
	})

	if len(parts) == 0 {
		return ""
	}

	if len(parts) == 1 {
		return parts[0]
	}

	parts = lo.Uniq(parts)

	parts = lo.Filter(parts, func(s string, index int) bool {
		s = strings.TrimSpace(s)
		if s == "" {
			return false
		}

		return true
	})

	if len(parts) == 1 {
		return parts[0]
	}

	return parts[0] + " (" + parts[1] + ")"
}

func SanitizeDisplayName(v string) string {
	if len(v) > 16 {
		v = v[:16]
	}

	bv := string(displayNameSanitizer.ReplaceAll([]byte(v), []byte("")))
	bv = strings.TrimSpace(v)

	return bv
}
