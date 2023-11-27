package webservice

import (
	"github.com/adrianrudnik/ablegram/internal/pusher"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"net/http"
)

func registerPusherRoutes(rg *gin.RouterGroup, pm *pusher.PushManager) {
	rg.GET("/clients", func(c *gin.Context) {
		if ok := isAdmin(c); !ok {
			return
		}

		// @todo Liste aller nutzer mit einer Liste aller verbindungen
		c.JSON(http.StatusOK, lo.Map(pm.GetClients(), func(client pusher.Client, _ int) pusher.UserClient {
			return pusher.UserClient{
				ClientId: client.ID,
				ClientIP: client.GetIP(true),
				UserId:   client.UserID,
			}
		}))
	})
}
