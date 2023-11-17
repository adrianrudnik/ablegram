package webservice

import (
	"github.com/adrianrudnik/ablegram/internal/config"
	"github.com/adrianrudnik/ablegram/internal/ui"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

func registerOsRoutes(rg *gin.RouterGroup, conf *config.Config) {
	rg.POST("/shutdown", func(c *gin.Context) {
		c.Status(202)
		c.JSON(200, gin.H{"message": "Shutting down"})

		if !conf.Behaviour.DemoMode {
			return
		}

		go func() {
			time.Sleep(500 * time.Millisecond)
			os.Exit(0)
		}()
	})

	rg.POST("/open", func(c *gin.Context) {
		var json OpenInput
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if !conf.Behaviour.DemoMode {
			ui.OpenDefault(json.Path)
		}

		c.JSON(200, gin.H{"status": "ok"})
	})
}
