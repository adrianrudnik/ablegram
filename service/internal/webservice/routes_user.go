package webservice

import (
	"github.com/adrianrudnik/ablegram/internal/access"
	"github.com/gin-gonic/gin"
	"net/http"
)

func registerUserRoutes(rg *gin.RouterGroup, users *access.UserList) {
	rg.GET("/users", func(c *gin.Context) {
		if ok := isAdmin(c); !ok {
			return
		}

		// @todo Liste aller nutzer mit einer Liste aller verbindungen
		c.JSON(http.StatusOK, users.All())
	})
}
