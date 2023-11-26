package webservice

import (
	"github.com/adrianrudnik/ablegram/internal/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func registerSuggestRoutes(rg *gin.RouterGroup, conf *config.Config) {
	rg.POST("/suggestion", func(c *gin.Context) {
		if ok := isSomeone(c); !ok {
			return
		}

		type userInput struct {
			ID string `json:"id" binding:"required"`
		}

		var input userInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if !conf.Behaviour.DemoMode {
			//ui.OpenDefault(input.Path)
		}

		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
}
