package webservice

import (
	"embed"
	"github.com/adrianrudnik/ablegram/internal/indexer"
	bleveHttp "github.com/blevesearch/bleve/v2/http"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/logger"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/rs/zerolog"
	"net/http"
	"os"
)

var Logger = zerolog.New(os.Stderr).With().Timestamp().Logger()

//go:embed .frontend/*
var frontendFs embed.FS

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Serve(indexer *indexer.Search, pushChannel *PushChannel, bindAddr string) error {
	// Wrap route logging into correct format
	// @see https://gin-gonic.com/docs/examples/define-format-for-the-log-of-routes/
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		Logger.Info().
			Str("method", httpMethod).
			Str("path", absolutePath).
			Str("handler", handlerName).
			Int("number-handlers", nuHandlers).
			Msg("Registering route")
	}

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(logger.SetLogger())

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:10000", "http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Accept-Encoding"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type", "Content-Encoding"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
	}))

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

	search := r.Group("/search")
	registerBleveRoutes(search, indexer)

	// Start a websocket server for UI related channels
	// @see https://medium.com/@abhishekranjandev/building-a-production-grade-websocket-for-notifications-with-golang-and-gin-a-detailed-guide-5b676dcfbd5a
	// @see https://github.com/tinkerbaj/chat-websocket-gin/blob/main/chat/chat.go

	r.GET("/ws", func(c *gin.Context) {
		connectClientWebsocket(c, pushChannel)
	})

	// Register the fallback route to the frontend UI bootstrap

	err = r.Run(bindAddr)
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

func registerBleveRoutes(rg *gin.RouterGroup, indexer *indexer.Search) {
	// @see https://github.com/blevesearch/beer-search/blob/master/main.go
	// @see https://github.com/blevesearch/beer-search/blob/master/http_util.go

	bleveHttp.RegisterIndexName("overview", indexer.Index)
	searchHandler := bleveHttp.NewSearchHandler("overview")
	rg.Any("/query", gin.WrapH(searchHandler))
}
