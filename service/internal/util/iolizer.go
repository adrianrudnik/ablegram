package util

import (
	"fmt"
	"github.com/samber/lo"
	"path/filepath"
	"runtime"
	"strings"
)

func PathContainsFolder(path string, folder string) bool {

	p := filepath.ToSlash(filepath.Dir(path))
	f := strings.Split(p, "/")

	return lo.Contains(f, folder)
}

func IsPathOriginFromTheSameOs(path string) bool {
	if len(path) == 0 {
		return false
	}

	// Most common scenarios:
	switch runtime.GOOS {
	case "windows":
		// We always have the drive letter
		if len(path) < 3 {
			return false
		}

		// We also always have a colon after the drive letter
		if fmt.Sprintf(string([]rune(path)[1])) != ":" {
			return false
		}

	case "darwin":
		// We always have a slash at the beginning
		if fmt.Sprintf(string([]rune(path)[0])) != "/" {
			return false
		}

		// Does it look like a home folder of linux?
		if strings.HasPrefix(path, "/home/") {
			return false
		}

	case "linux":
		// We always have a slash at the beginning
		if fmt.Sprintf(string([]rune(path)[0])) != "/" {
			return false
		}

		// Does it look like a home folder of mac?
		if strings.HasPrefix(path, "/Users/") {
			return false
		}
	}

	return true
}
