package pusher

import (
	"github.com/gin-gonic/gin"
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

func (c *PushChannel) ConnectClientWebsocket(ctx *gin.Context) {
	ws, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		Logger.Error().Err(err).Msg("Failed to upgrade client to websocket")
		return
	}
	client := NewPushClient(ws, c)

	client.DisplayName = ctx.GetString("displayName")
	client.Role = ctx.GetString("role")

	c.addClient <- client

	go client.Send()
	go client.Receive()
}
