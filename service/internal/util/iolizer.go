package util

import (
	"github.com/samber/lo"
	"path/filepath"
	"strings"
)

func PathContainsFolder(path string, folder string) bool {

	p := filepath.ToSlash(filepath.Dir(path))
	f := strings.Split(p, "/")

	return lo.Contains(f, folder)
}
