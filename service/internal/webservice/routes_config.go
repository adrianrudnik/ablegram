package webservice

import (
	"github.com/adrianrudnik/ablegram/internal/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func registerConfigRoutes(rg *gin.RouterGroup, conf *config.Config) {
	rg.GET("/config", func(c *gin.Context) {
		c.JSON(200, conf)
	})

	rg.PUT("/config/log", func(c *gin.Context) {
		type userInput struct {
			Level                  string `json:"level"`
			EnableRuntimeLogfile   bool   `json:"enable_runtime_logfile"`
			EnableProcessedLogfile bool   `json:"enable_processed_logfile"`
		}

		var input userInput

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Early exit in demo mode, we do not want to save anything
		if conf.Behaviour.DemoMode {
			c.JSON(200, conf)
			return
		}

		conf.Log.Level = input.Level

		if input.EnableRuntimeLogfile {
			conf.Log.EnableRuntimeLogfile = true
			conf.Log.EnableProcessedLogfile = input.EnableProcessedLogfile
		} else {
			conf.Log.EnableProcessedLogfile = false
		}

		err := conf.Save()
		if err != nil {
			Logger.Error().Err(err).Msg("Could not save configuration")
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, conf)
	})

	rg.PUT("/config/behaviour", func(c *gin.Context) {
		type userInput struct {
			AutostartWebservice bool `json:"autostart_webservice"`
			OpenBrowserOnStart  bool `json:"open_browser_on_start"`
			ShowServiceGui      bool `json:"show_service_gui"`
		}

		var input userInput

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Early exit in demo mode, we do not want to save anything
		if conf.Behaviour.DemoMode {
			c.JSON(200, conf)
			return
		}

		conf.Behaviour.AutostartWebservice = input.AutostartWebservice
		conf.Behaviour.OpenBrowserOnStart = input.OpenBrowserOnStart
		conf.Behaviour.ShowServiceGui = input.ShowServiceGui

		err := conf.Save()
		if err != nil {
			Logger.Error().Err(err).Msg("Could not save configuration")
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, conf)
	})

	// Register new collector
	rg.POST("/config/collector", func(c *gin.Context) {
		type userInput struct {
			ID                   string `json:"id"`
			Type                 string `json:"type"`
			Uri                  string `json:"uri"`
			ParserPerformance    string `json:"parser_performance"`
			ParserWorkerDelay    int    `json:"parser_delay"`
			ExcludeSystemFolders bool   `json:"exclude_system_folders"`
			ExcludeDotFolders    bool   `json:"exclude_dot_folders"`
		}

		var input userInput

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Early exit in demo mode, we do not want to save anything
		if conf.Behaviour.DemoMode {
			c.JSON(200, conf)
			return
		}

		conf.Collector.Targets = append(conf.Collector.Targets, config.CollectorTarget{
			ID:                   input.ID,
			Uri:                  input.Uri,
			ParserPerformance:    input.ParserPerformance,
			ParserWorkerDelay:    input.ParserWorkerDelay,
			ExcludeSystemFolders: input.ExcludeSystemFolders,
			ExcludeDotFolders:    input.ExcludeDotFolders,
		})

		err := conf.Save()
		if err != nil {
			Logger.Error().Err(err).Msg("Could not save configuration")
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, conf)
	})

	// Update existing collector
	rg.PUT("/config/collector/:id", func(c *gin.Context) {
		id := c.Param("id")

		type userInput struct {
			Uri                  string `json:"uri"`
			ParserPerformance    string `json:"parser_performance"`
			ParserWorkerDelay    int    `json:"parser_delay"`
			ExcludeSystemFolders bool   `json:"exclude_system_folders"`
			ExcludeDotFolders    bool   `json:"exclude_dot_folders"`
		}

		var input userInput

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Early exit in demo mode, we do not want to save anything
		if conf.Behaviour.DemoMode {
			c.JSON(200, conf)
			return
		}

		for i, target := range conf.Collector.Targets {
			if target.ID == id {
				conf.Collector.Targets[i].Uri = input.Uri
				conf.Collector.Targets[i].ParserPerformance = input.ParserPerformance
				conf.Collector.Targets[i].ParserWorkerDelay = input.ParserWorkerDelay
				conf.Collector.Targets[i].ExcludeSystemFolders = input.ExcludeSystemFolders
				conf.Collector.Targets[i].ExcludeDotFolders = input.ExcludeDotFolders
				break
			}
		}

		err := conf.Save()
		if err != nil {
			Logger.Error().Err(err).Msg("Could not save configuration")
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, conf)
	})

	// Delete collector
	rg.DELETE("/config/collector/:id", func(c *gin.Context) {
		id := c.Param("id")

		// Early exit in demo mode, we do not want to save anything
		if conf.Behaviour.DemoMode {
			c.JSON(200, conf)
			return
		}

		for i, target := range conf.Collector.Targets {
			if target.ID == id {
				conf.Collector.Targets = append(conf.Collector.Targets[:i], conf.Collector.Targets[i+1:]...)
				break
			}
		}

		err := conf.Save()
		if err != nil {
			Logger.Error().Err(err).Msg("Could not save configuration")
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, conf)
	})
}
