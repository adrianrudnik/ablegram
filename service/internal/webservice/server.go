package webservice

import (
	"embed"
	"github.com/adrianrudnik/ablegram/internal/access"
	"github.com/adrianrudnik/ablegram/internal/config"
	"github.com/adrianrudnik/ablegram/internal/indexer"
	"github.com/adrianrudnik/ablegram/internal/pusher"
	"github.com/adrianrudnik/ablegram/internal/tagger"
	"github.com/adrianrudnik/ablegram/internal/ui"
	"github.com/adrianrudnik/ablegram/internal/workload"
	bleveHttp "github.com/blevesearch/bleve/v2/http"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

//go:embed .frontend/*
var frontendFs embed.FS

func Serve(conf *config.Config, auth *access.Auth, otp *access.Otp, indexer *indexer.Search, tc *tagger.TagCollector, pushChan chan workload.PushMessage, bindAddr string) error {
	// Wrap route logging into correct format
	// @see https://gin-gonic.com/docs/examples/define-format-for-the-log-of-routes/
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		Logger.Debug().
			Str("method", httpMethod).
			Str("path", absolutePath).
			Str("handler", handlerName).
			Int("number-handlers", nuHandlers).
			Msg("Registering route")
	}

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(LoggerMiddleware())
	r.Use(CacheMiddleware())
	r.Use(AccessMiddleware(auth))

	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Accept-Encoding"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
	}))

	// Set the trusted platform
	// @see https://github.com/gin-gonic/gin/blob/44d0dd70924dd154e3b98bc340accc53484efa9c/gin.go#L73C1-L80C2
	r.TrustedPlatform = conf.Webservice.TrustedPlatform

	pusher.Logger = Logger.With().Str("module", "pusher").Logger()
	pushManager := pusher.NewPushManager(conf, pushChan)
	go pushManager.Run()

	// Mount the embeded search frontend
	frontendFS := EmbedFolder(frontendFs, ".frontend")

	// Mount the Vue frontend
	r.Use(static.Serve("/", frontendFS))

	// Mount a fallback to the Vue frontend, if no matching route was found
	r.NoRoute(func(c *gin.Context) {
		c.FileFromFS("/", frontendFS)
	})

	// Start a websocket server for UI related channels
	// @see https://medium.com/@abhishekranjandev/building-a-production-grade-websocket-for-notifications-with-golang-and-gin-a-detailed-guide-5b676dcfbd5a
	// @see https://github.com/tinkerbaj/chat-websocket-gin/blob/main/chat/chat.go

	r.GET("/ws", func(c *gin.Context) {
		pushManager.ConnectClientWebsocket(c)
	})

	// Register common API routes
	api := r.Group("/api")
	registerApiRoutes(api)
	registerTagRoutes(api, tc)
	registerConfigRoutes(api, conf)
	registerOsRoutes(api, conf)

	// Boot up auth and otp services
	registerAccessRoutes(api, conf, auth)

	// Register the bleve HTTP router
	search := r.Group("/search")
	registerBleveRoutes(search, indexer)

	r.GET("/about", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"version": conf.About.Version,
			"commit":  conf.About.Commit,
			"date":    conf.About.Date,
		})
	})

	// Issue a single OTP and admin URL to the service console when we are in dev mode.
	if conf.IsDevelopmentEnv {
		log.Info().
			Str("url", ui.GenerateLocalAdminUrl(conf, otp)).
			Msg("Generated Admin OTP url")
	}

	// Register the fallback route to the frontend UI bootstrap
	err := r.Run(bindAddr)
	if err != nil {
		return err
	}

	return nil
}

func registerTagRoutes(rg *gin.RouterGroup, tc *tagger.TagCollector) {
	rg.GET("/tags", func(c *gin.Context) {
		if c.Query("verbose") != "" {
			c.IndentedJSON(200, tc.GetDetailedTags())
		} else {
			c.IndentedJSON(200, tc.GetBaseTags())
		}
	})
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
