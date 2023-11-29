package pusher

import (
	"github.com/adrianrudnik/ablegram/internal/util"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"net"
	"time"
)

// Client represents a single connection to the system.
type Client struct {
	ID     uuid.UUID
	UserID uuid.UUID

	Conn     *websocket.Conn
	Tx       chan interface{}
	pushChan *PushManager
}

// UserClient represents a user that is currently connected to the system.
// A client can have multiple connections (e.g. multiple tabs).
type UserClient struct {
	ClientId uuid.UUID `json:"id"`
	ClientIP net.IP    `json:"ip"`
	UserId   uuid.UUID `json:"user_id"`
}

const writeTimeout = 10 * time.Second
const pongWait = 15 * time.Second
const pingPeriod = (pongWait * 9) / 10

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
	pingTicker := time.NewTicker(pingPeriod)

	defer func() {
		pingTicker.Stop()

		err := c.Conn.Close()
		Logger.Debug().Err(err).Msg("Ending client tx routing")
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
				Logger.Warn().Err(err).Msg("Could not set write deadline for client on ping tick")
			}

			if err := c.Conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				Logger.Debug().Err(err).Msg("Could not send ping message to client")
				return
			}
		}
	}
}

func (c *Client) Receive() {
	// Ensure a failing routine cleans up the client
	defer func() {
		c.pushChan.removeClient <- c
		err := c.Conn.Close()
		Logger.Debug().Err(err).Msg("Ending client rx routing")
	}()

	c.Conn.SetReadLimit(512)
	err := c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	if err != nil {
		Logger.Warn().Err(err).Msg("Failed to set read deadline on websocket")
		return
	}

	c.Conn.SetPongHandler(func(string) error {
		err := c.Conn.SetReadDeadline(time.Now().Add(pongWait))
		if err != nil {
			Logger.Warn().Err(err).Msg("Failed to set read deadline in pong handler")
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
