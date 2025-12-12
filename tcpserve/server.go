package tcpserve

import (
	"context"
	"fmt"
	"github.com/chilledoj/sockt"
	"github.com/chilledoj/sockt-examples/types"
	"net"
	"sync/atomic"
)

var playerCounter atomic.Int32

type TcpServer struct {
	listener *net.TCPListener
	room     *sockt.Room[types.RoomID, types.PlayerID]
}

func NewTcpServer(room *sockt.Room[types.RoomID, types.PlayerID]) (*TcpServer, error) {
	// Resolve the string address to a TCP address
	tcpAddr, err := net.ResolveTCPAddr("tcp4", ":9191")

	if err != nil {
		return nil, err
	}

	// Start listening for TCP connections on the given address
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		return nil, err
	}

	return &TcpServer{listener: listener, room: room}, nil
}

func (t *TcpServer) Start() error {
	for {
		// Accept new connections
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Println(err)
			return err
		}

		c := &NetConnWrapper{Conn: conn}
		playNum := playerCounter.Add(1)
		playerID := fmt.Sprintf("player%02d", playNum)

		if err := t.room.AddConnection(c, playerID); err != nil {
			conn.Close()
			fmt.Println(err)
		}
	}
	return nil
}

func (t *TcpServer) Stop(ctx context.Context) error {
	return t.listener.Close()
}
