package webservice

import (
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"path/filepath"
)

func CacheControl() gin.HandlerFunc {
	return func(c *gin.Context) {
		cacheableSuffixes := []string{
			".css",
			".js",
			".png",
			".jpg",
			".jpeg",
			".gif",
			".ico",
			".svg",
			".woff",
			".woff2",
			".ttf",
			".eot",
			".otf",
		}

		_, found := lo.Find(cacheableSuffixes, func(prefix string) bool {
			return filepath.Ext(c.Request.RequestURI) == prefix
		})

		if found {
			c.Writer.Header().Set("Cache-Control", "public, max-age=604800, immutable")
		} else {
			c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		}
	}
}
