package webservice

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"time"
)
import "github.com/google/uuid"

type PushClient struct {
	ID         string
	Connection *websocket.Conn
	tx         chan interface{}
	channel    *PushChannel
}

var writeTimeout = 5 * time.Second

func NewPushClient(connection *websocket.Conn, channel *PushChannel) *PushClient {
	id := uuid.New()

	return &PushClient{
		ID:         id.String(),
		Connection: connection,

		channel: channel,
		tx:      make(chan interface{}, 128),
	}
}

func (c *PushClient) Send() {
	// Set up a ticker to produce the ping message in intervals
	pingTicker := time.NewTicker(5 * time.Second)

	defer func() {
		pingTicker.Stop()
		c.Connection.Close()
	}()

	for {
		select {
		case msg, ok := <-c.tx:
			err := c.Connection.SetWriteDeadline(time.Now().Add(writeTimeout))
			if err != nil {
				Logger.Warn().Err(err).Msg("Could not set write deadline for client on message")
			}

			if !ok {
				err := c.Connection.WriteMessage(websocket.CloseMessage, []byte{})
				if err != nil {
					Logger.Warn().Err(err).Msg("Could not send close message to client")
				}
				return
			}

			err = c.Connection.WriteJSON(msg)
			if err != nil {
				Logger.Error().Err(err).Msg("Could not write JSON message to client")
			}

		case <-pingTicker.C:
			err := c.Connection.SetWriteDeadline(time.Now().Add(writeTimeout))
			if err != nil {
				Logger.Warn().Err(err).Msg("Could not set write deadline for client on ping")
			}

			err = c.Connection.WriteMessage(websocket.PingMessage, nil)
			if err != nil {
				Logger.Warn().Err(err).Msg("Could not write ping message to client")
			}
		}
	}
}

func (c *PushClient) Receive() {
	// Ensure a failing routine cleans up the client
	defer func() {
		c.channel.removeClient <- c
		c.Connection.Close()
	}()

	c.Connection.SetPongHandler(func(string) error {
		err := c.Connection.SetReadDeadline(time.Now().Add(1 * time.Minute))
		if err != nil {
			return err
		}
		return nil
	})

	for {
		// Currently we do not support client messages, so we just keep the connection empty
		_, _, err := c.Connection.ReadMessage()
		if err != nil {
			Logger.Error().Err(err).Msg("Failed to read client message")
			break
		}
	}
}

func connectClientWebsocket(ctx *gin.Context, channel *PushChannel) {
	ws, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		Logger.Error().Err(err).Msg("Failed to upgrade client to websocket")
		return
	}

	c := NewPushClient(ws, channel)

	channel.addClient <- c

	go c.Send()
	go c.Receive()
}
