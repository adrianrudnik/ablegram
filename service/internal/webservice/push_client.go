package webservice

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"time"
)
import "github.com/google/uuid"

type PushClient struct {
	ID       string
	Conn     *websocket.Conn
	tx       chan interface{}
	pushChan *PushChannel
}

var writeTimeout = 5 * time.Second

func NewPushClient(connection *websocket.Conn, pushChan *PushChannel) *PushClient {
	id := uuid.New()

	return &PushClient{
		ID:   id.String(),
		Conn: connection,

		pushChan: pushChan,
		tx:       make(chan interface{}, 128),
	}
}

func (c *PushClient) Send() {
	// Set up a ticker to produce the ping message in intervals
	pingTicker := time.NewTicker(5 * time.Second)

	defer func() {
		pingTicker.Stop()
		c.Conn.Close()
	}()

	for {
		select {
		case msg, ok := <-c.tx:
			err := c.Conn.SetWriteDeadline(time.Now().Add(writeTimeout))
			if err != nil {
				Logger.Warn().Err(err).Msg("Could not set write deadline for client on message")
			}

			if !ok {
				err := c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				if err != nil {
					Logger.Info().Err(err).Msg("Could not send close message to client")
				}
				return
			}

			err = c.Conn.WriteJSON(msg)
			if err != nil {
				Logger.Error().Err(err).Msg("Could not write JSON message to client")
			}

		case <-pingTicker.C:
			err := c.Conn.SetWriteDeadline(time.Now().Add(writeTimeout))
			if err != nil {
				Logger.Warn().Err(err).Msg("Could not set write deadline for client on ping")
			}

			err = c.Conn.WriteMessage(websocket.PingMessage, nil)
			if err != nil {
				Logger.Warn().Err(err).Msg("Could not write ping message to client")
			}
		}
	}
}

func (c *PushClient) Receive() {
	// Ensure a failing routine cleans up the client
	defer func() {
		c.pushChan.removeClient <- c
		c.Conn.Close()
	}()

	c.Conn.SetPongHandler(func(string) error {
		err := c.Conn.SetReadDeadline(time.Now().Add(1 * time.Minute))
		if err != nil {
			return err
		}
		return nil
	})

	for {
		// Currently we do not support client messages, so we just keep the connection empty
		_, _, err := c.Conn.ReadMessage()
		if err != nil {
			Logger.Info().Err(err).Msg("Failed to read client message")
			break
		}
	}
}

func connectClientWebsocket(ctx *gin.Context, pushChan *PushChannel) {
	ws, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		Logger.Error().Err(err).Msg("Failed to upgrade client to websocket")
		return
	}

	c := NewPushClient(ws, pushChan)

	pushChan.addClient <- c

	go c.Send()
	go c.Receive()
}
