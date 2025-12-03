package gorillaws

import (
	"context"
	"github.com/chilledoj/sockt"
	"github.com/gorilla/websocket"
)

type GorillaSocketWrapper struct {
	Conn *websocket.Conn
}

func (c *GorillaSocketWrapper) Write(messageType sockt.SocketMessageType, bytes []byte) error {
	csType := websocket.TextMessage
	if messageType == sockt.SocketMessageBinary {
		csType = websocket.BinaryMessage
	}
	return c.Conn.WriteMessage(csType, bytes)
}

func (c *GorillaSocketWrapper) Read(ctx context.Context) ([]byte, error) {
	_, bytes, err := c.Conn.ReadMessage()

	return bytes, err
}
