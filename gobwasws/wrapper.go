package gobwasws

import (
	"context"
	"github.com/chilledoj/sockt"
	"github.com/gobwas/ws/wsutil"
	"net"
)

var _ sockt.Socket = GobwasWsSocketWrapper{}

type GobwasWsSocketWrapper struct {
	Conn net.Conn
}

func (c GobwasWsSocketWrapper) Close() error {
	return c.Conn.Close()
}

func (c GobwasWsSocketWrapper) Write(messageType sockt.SocketMessageType, bytes []byte) error {
	return wsutil.WriteServerBinary(c.Conn, bytes)
}

func (c GobwasWsSocketWrapper) Read(ctx context.Context) ([]byte, error) {
	bytes, _, err := wsutil.ReadClientData(c.Conn)

	return bytes, err
}
