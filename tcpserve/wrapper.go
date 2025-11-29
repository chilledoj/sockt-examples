package tcpserve

import (
	"bufio"
	"context"
	"fmt"
	"github.com/chilledoj/sockt"
	"net"
)

type NetConnWrapper struct {
	Conn net.Conn
}

func (c *NetConnWrapper) Write(messageType sockt.SocketMessageType, bytes []byte) error {
	_, err := c.Conn.Write(bytes)
	return err
}

func (c *NetConnWrapper) Read(ctx context.Context) ([]byte, error) {
	buf := make([]byte, 4096)

	n, err := bufio.NewReader(c.Conn).Read(buf)
	if err != nil {
		return nil, err
	}
	if n == 0 {
		return nil, fmt.Errorf("no data read")
	}
	return buf[:n], nil
}
