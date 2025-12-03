package gorillaws

import (
	"context"
	"fmt"
	"github.com/chilledoj/sockt"
	"github.com/chilledoj/sockt-examples/types"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync/atomic"
)

var playerCounter atomic.Int32

type GorillaServer struct {
	s *http.Server
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func NewGorillaServer(room *sockt.Room[types.RoomID, types.PlayerID]) *GorillaServer {

	mux := http.NewServeMux()

	s := &http.Server{
		Addr:    ":9090",
		Handler: mux,
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`<html><body><h1>Welcome to the Coder's Room!</h1><a href="/api/ws">Connect to the WebSocket</a></body></html>`))
	})

	mux.HandleFunc("/api/ws", func(w http.ResponseWriter, r *http.Request) {
		// pretend to get playerID + Room ID
		playNum := playerCounter.Add(1)
		playerID := fmt.Sprintf("player%02d", playNum)

		//var routeRoomID RoomID = roomID

		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}

		conn := &GorillaSocketWrapper{Conn: c}

		room.AddConnection(conn, playerID)

		// Keep the handler alive until context is done or connection closes
		//<-r.Context().Done()
	})

	return &GorillaServer{s: s}
}

func (c *GorillaServer) Start() error {
	return c.s.ListenAndServe()
}

func (c *GorillaServer) Stop(ctx context.Context) error {
	return c.s.Shutdown(ctx)
}
