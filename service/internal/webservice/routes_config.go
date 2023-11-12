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
}
