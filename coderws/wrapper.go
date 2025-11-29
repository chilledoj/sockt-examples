package coderws

import (
	"context"
	"fmt"
	"github.com/chilledoj/sockt"
	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
)

type CoderSocketWrapper struct {
	Conn *websocket.Conn
}

func (c *CoderSocketWrapper) Write(messageType sockt.SocketMessageType, bytes []byte) error {
	csType := websocket.MessageText
	if messageType == sockt.SocketMessageBinary {
		csType = websocket.MessageBinary
	}
	return c.Conn.Write(context.Background(), csType, bytes)
}

func (c *CoderSocketWrapper) Read(ctx context.Context) ([]byte, error) {
	var r interface{}
	err := wsjson.Read(ctx, c.Conn, &r)
	return []byte(fmt.Sprintf("%v", r)), err
}
