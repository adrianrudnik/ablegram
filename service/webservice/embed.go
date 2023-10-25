package webservice

import (
	"embed"
	"github.com/gin-contrib/static"
	"io/fs"
	"net/http"
)

// @see https://github.com/gin-contrib/static/issues/19

type embedFS struct {
	http.FileSystem
}

func (fs embedFS) Exists(prefix string, path string) bool {
	_, err := fs.Open(path)
	if err != nil {
		return false
	}
	return true
}

func EmbedFolder(fsEmbed embed.FS, targetPath string) static.ServeFileSystem {
	f, err := fs.Sub(fsEmbed, targetPath)
	if err != nil {
		panic(err)
	}
	return embedFS{
		FileSystem: http.FS(f),
	}
}
