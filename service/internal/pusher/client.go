package pusher

import (
	"github.com/adrianrudnik/ablegram/internal/util"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"net"
	"time"
)

type Client struct {
	ID     uuid.UUID
	UserID uuid.UUID

	Conn     *websocket.Conn
	Tx       chan interface{}
	pushChan *PushManager
}

var writeTimeout = 5 * time.Second

func NewClient(id uuid.UUID, userId uuid.UUID, connection *websocket.Conn, pushChan *PushManager) *Client {
	return &Client{
		ID:     id,
		UserID: userId,

		Conn: connection,

		pushChan: pushChan,
		Tx:       make(chan interface{}, 128),
	}
}

func (c *Client) Send() {
	// Set up a ticker to produce the ping message in intervals
	pingTicker := time.NewTicker(5 * time.Second)

	defer func() {
		pingTicker.Stop()
		c.Conn.Close()
	}()

	for {
		select {
		case msg, ok := <-c.Tx:
			err := c.Conn.SetWriteDeadline(time.Now().Add(writeTimeout))
			if err != nil {
				Logger.Warn().Err(err).Msg("Could not set write deadline for client on message")
			}

			if !ok {
				err := c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				if err != nil {
					Logger.Debug().Err(err).Msg("Could not send close message to client")
				}
				return
			}

			if !c.pushChan.canClientReceive(msg, c) {
				continue
			}

			// We made it here, we can broadcast it to the client
			err = c.Conn.WriteJSON(msg)
			if err != nil {
				Logger.Debug().Err(err).Msg("Could not write JSON message to client")
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

func (c *Client) Receive() {
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
			Logger.Debug().Err(err).Msg("Failed to read client message")
			break
		}
	}
}

func (c *Client) GetIP(real bool) net.IP {
	if !real {
		return util.GetFakeClientIP()
	}

	return c.Conn.RemoteAddr().(*net.TCPAddr).IP
}
