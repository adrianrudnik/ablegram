package apiserver

import (
	"embed"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

//go:embed .frontend/*
var frontendFs embed.FS

func Serve(bind string) error {
	r := gin.Default()

	// No reason to support proxies yet
	err := r.SetTrustedProxies([]string{})
	if err != nil {
		return err
	}

	// Mount the Vue frontend
	r.Use(static.Serve("/", EmbedFolder(frontendFs, ".frontend")))

	api := r.Group("/api")

	registerApiRoutes(api)

	err = r.Run(bind)
	if err != nil {
		return err
	}

	return nil
}

func registerApiRoutes(rg *gin.RouterGroup) {
	rg.GET("/status", func(c *gin.Context) {
		c.String(200, "pong")
	})
}
