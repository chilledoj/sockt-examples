package gobwasws

import (
	"context"
	"fmt"
	"github.com/chilledoj/sockt"
	"github.com/chilledoj/sockt-examples/types"
	"github.com/gobwas/ws"
	"log"
	"net/http"
	"sync/atomic"
)

var playerCounter atomic.Int32

type GobwasWSServer struct {
	s *http.Server
}

func NewGobwasWSServer(room *sockt.Room[types.RoomID, types.PlayerID]) *GobwasWSServer {

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

		c, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
			log.Println(err)
			return
		}

		conn := &GobwasWsSocketWrapper{Conn: c}

		if err := room.AddConnection(conn, playerID); err != nil {
			c.Close()
			http.Error(w, err.Error(), http.StatusBadRequest)

			return
		}
		// Keep connection alive
		<-r.Context().Done()

	})

	return &GobwasWSServer{s: s}
}

func (c *GobwasWSServer) Start() error {
	return c.s.ListenAndServe()
}

func (c *GobwasWSServer) Stop(ctx context.Context) error {
	return c.s.Shutdown(ctx)
}
