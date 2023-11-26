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
		if ok := isAdmin(c); !ok {
			return
		}

		c.JSON(http.StatusAccepted, gin.H{"message": "Shutting down"})

		if conf.Behaviour.DemoMode {
			return
		}

		go func() {
			time.Sleep(500 * time.Millisecond)
			os.Exit(0)
		}()
	})

	rg.POST("/open", func(c *gin.Context) {
		if ok := isAdmin(c); !ok {
			return
		}

		type userInput struct {
			Path string `json:"path" binding:"required"`
		}

		var input userInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if !conf.Behaviour.DemoMode {
			ui.OpenDefault(input.Path)
		}

		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
}
