package webservice

import (
	"github.com/adrianrudnik/ablegram/internal/config"
	"github.com/adrianrudnik/ablegram/internal/suggest"
	"github.com/gin-gonic/gin"
	"net/http"
)

func registerSuggestRoutes(rg *gin.RouterGroup, conf *config.Config, suggest *suggest.List) {
	rg.GET("/suggestions", func(c *gin.Context) {
		if ok := isAdmin(c); !ok {
			return
		}

		c.JSON(http.StatusOK, suggest.All())
	})

	rg.POST("/suggestions", func(c *gin.Context) {
		// All users that said hello can suggest
		if ok := isSomeone(c); !ok {
			return
		}

		type userInput struct {
			Target string `json:"target" binding:"required"`
		}

		var input userInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
}
