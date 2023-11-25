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

func ConnectClientWebsocket(ctx *gin.Context, pushChan *PushChannel) {
	ws, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		Logger.Error().Err(err).Msg("Failed to upgrade client to websocket")
		return
	}

	c := NewPushClient(ws, pushChan)

	c.DisplayName = ctx.GetString("displayName")
	c.Role = ctx.GetString("role")

	pushChan.addClient <- c

	go c.Send()
	go c.Receive()
}
