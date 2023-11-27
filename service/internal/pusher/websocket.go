package pusher

import (
	"github.com/adrianrudnik/ablegram/internal/auth"
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

func (pm *PushManager) ConnectClientWebsocket(ctx *gin.Context) {
	// Generate a unique client ID and communicate it back to the client.
	// Same user might connect multiple times (e.g. multiple tabs), so we need to distinguish them.
	clientId := uuid.New()

	ws, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		Logger.Error().Err(err).Msg("Failed to upgrade client to websocket")
		return
	}

	client := NewClient(clientId, ctx.MustGet("userId").(uuid.UUID), ws, pm)
	client.UserID = ctx.MustGet("userId").(uuid.UUID)

	// Add the user to the known list
	pm.users.Add(auth.NewUser(client.UserID, ctx.GetString("userDisplayName"), ctx.GetString("userRole")))

	pm.addClient <- client

	go client.Send()
	go client.Receive()
}
