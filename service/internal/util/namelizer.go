package util

import (
	"github.com/samber/lo"
	"strings"
)

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
