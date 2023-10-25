package apiserver

import (
	"embed"
	"github.com/gin-contrib/logger"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"os"
)
import "github.com/gorilla/websocket"

var Logger = zerolog.New(os.Stderr).With().Timestamp().Logger()

//go:embed .frontend/*
var frontendFs embed.FS

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Serve(bind string) error {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(logger.SetLogger())

	// No reason to support proxies yet
	err := r.SetTrustedProxies([]string{})
	if err != nil {
		return err
	}

	frontendFS := EmbedFolder(frontendFs, ".frontend")

	// Mount the Vue frontend
	r.Use(static.Serve("/", frontendFS))

	// Mount a fallback to the Vue frontend, if no matching route was found
	r.NoRoute(func(c *gin.Context) {
		c.Status(200)
		c.FileFromFS("index.html", frontendFS)
	})

	api := r.Group("/api")
	registerApiRoutes(api)

	// Start a websocket server for UI related channels
	// @see https://medium.com/@abhishekranjandev/building-a-production-grade-websocket-for-notifications-with-golang-and-gin-a-detailed-guide-5b676dcfbd5a
	// @see https://github.com/tinkerbaj/chat-websocket-gin/blob/main/chat/chat.go

	r.GET("/ws", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			return
		}
		defer conn.Close()
	})

	// Register the fallback route to the frontend UI bootstrap

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
