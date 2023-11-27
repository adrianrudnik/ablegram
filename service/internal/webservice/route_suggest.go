package webservice

import (
	"github.com/adrianrudnik/ablegram/internal/config"
	"github.com/adrianrudnik/ablegram/internal/suggest"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func registerSuggestRoutes(rg *gin.RouterGroup, conf *config.Config, sl *suggest.List) {
	rg.GET("/remarks", func(c *gin.Context) {
		if ok := isAdmin(c); !ok {
			return
		}

		c.JSON(http.StatusOK, sl.All())
	})

	rg.POST("/remarks", func(c *gin.Context) {
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

	rg.DELETE("/remarks/:id", func(c *gin.Context) {
		if ok := isAdmin(c); !ok {
			return
		}

		sl.Delete(uuid.MustParse(c.Param("id")))

		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
}
