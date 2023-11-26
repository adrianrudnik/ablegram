package pusher

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (c *PushManager) ConnectClientWebsocket(ctx *gin.Context) {
	// Generate a unique client ID and communicate it back to the client.
	clientId := uuid.New()

	ws, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		Logger.Error().Err(err).Msg("Failed to upgrade client to websocket")
		return
	}

	client := NewPushClient(clientId, ws, c)

	client.DisplayName = ctx.GetString("userDisplayName")
	client.Role = ctx.GetString("userRole")

	c.addClient <- client

	go client.Send()
	go client.Receive()
}
