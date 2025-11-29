package coderws

import (
	"context"
	"fmt"
	"github.com/chilledoj/sockt"
	"github.com/chilledoj/sockt-examples/types"
	"github.com/coder/websocket"
	"net/http"
	"sync/atomic"
)

var playerCounter atomic.Int32

type CoderServer struct {
	s *http.Server
}

func NewCoderServer(room *sockt.Room[types.RoomID, types.PlayerID]) *CoderServer {

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

		c, err := websocket.Accept(w, r, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer c.CloseNow()

		conn := &CoderSocketWrapper{Conn: c}

		room.AddConnection(conn, playerID)
		// Keep the handler alive until context is done or connection closes
		<-r.Context().Done()
	})

	return &CoderServer{s: s}
}

func (c *CoderServer) Start() error {
	return c.s.ListenAndServe()
}

func (c *CoderServer) Stop(ctx context.Context) error {
	return c.s.Shutdown(ctx)
}
