package util

import (
	"strings"
	"unicode"
)

func EvaluateUserInput(value string) (safe string, empty bool) {
	// Remove everything that is not safe for human eyes
	safe = strings.Map(func(r rune) rune {
		if unicode.IsPrint(r) {
			return r
		}
		return -1
	}, value)

	safe = strings.TrimSpace(value)
	empty = len([]rune(safe)) == 0

	return
}
