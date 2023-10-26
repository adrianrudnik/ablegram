package util

import (
	"github.com/samber/lo"
	"strings"
)

func Namelize(parts []string) string {
	if len(parts) == 0 {
		return "[empty]"
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
