package engine

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/chilledoj/sockt"
	"github.com/chilledoj/sockt-examples/types"
	"log"
	"os"
	"sync"
	"time"
)

type Engine[PID comparable] struct {
	ctx    context.Context
	cancel context.CancelFunc

	// data
	mu      sync.RWMutex
	players []PID

	// Messaging
	sendChannel chan<- sockt.Event[PID]

	// Logging
	lg *log.Logger
}

func NewEngine[PID comparable](ctx context.Context) *Engine[PID] {
	ctx, cancel := context.WithCancel(ctx)

	return &Engine[PID]{
		ctx:     ctx,
		cancel:  cancel,
		players: make([]PID, 0),
		lg:      log.New(os.Stdout, "[engine] ", log.LstdFlags),
	}
}

func (e *Engine[PID]) Start() {
	tick := time.NewTicker(2 * time.Second)
	defer tick.Stop()
	for {
		select {
		case <-tick.C:
			e.debugLog()
		case <-e.ctx.Done():
			return
		}
	}
}

func (e *Engine[PID]) debugLog() {
	e.mu.RLock()
	defer e.mu.RUnlock()
	e.lg.Printf("Players: %d\n", len(e.players))
	for _, pid := range e.players {
		e.lg.Printf("Player: %v\n", pid)
	}
}

func (e *Engine[PID]) Stop() {
	e.cancel()
}

func (e *Engine[PID]) Init(c chan<- sockt.Event[PID]) {
	e.sendChannel = c
	go e.Start()
}

func (e *Engine[PID]) Process(msg sockt.Event[PID]) {
	e.lg.Printf("PROCESS: %v\n", msg)
	switch msg.Type {
	case sockt.EventConnect:
		e.processNewConnection(msg.Subject)
	case sockt.EventMessage:
		e.processMessage(msg)
	case sockt.EventDisconnect:
		e.processDisconnection(msg.Subject)
	}
}

func (e *Engine[PID]) processNewConnection(connID PID) {
	e.mu.Lock()
	e.players = append(e.players, connID)
	e.mu.Unlock()
	e.lg.Printf("Player connected: %v\n", connID)

	// Welcome Message
	e.lg.Printf("Sending welcome message to %v\n", connID)
	e.sendChannel <- sockt.Event[PID]{
		Type:    sockt.EventMessage,
		Subject: connID,
		Data:    []byte(`{"message": "Welcome to the Coder's Room!"}`),
	}

	// Notify other users of joining player
	e.notifyPlayers(connID, "join")
}

func (e *Engine[PID]) processMessage(msg sockt.Event[PID]) {
	var m types.Msg
	if err := json.Unmarshal(msg.Data, &m); err != nil {
		e.lg.Printf("Error unmarshalling message: %v\n", err)
	}

	e.lg.Printf("Message from %v: %v\n", msg.Subject, m)

}

func (e *Engine[PID]) processDisconnection(connID PID) {
	e.mu.Lock()

	for i, pid := range e.players {
		if pid == connID {
			e.players = append(e.players[:i], e.players[i+1:]...)
			break
		}
	}
	e.mu.Unlock()
	e.lg.Printf("Player disconnected: %v\n", connID)
	// Notify other users of joining player
	e.notifyPlayers(connID, "depart")
}

func (e *Engine[PID]) notifyPlayers(excludePid PID, action string) {
	e.lg.Printf("Notifying other players of player %v %sing\n", excludePid, action)
	if len(e.players) == 0 {
		return
	}
	e.mu.RLock()
	recipients := make([]PID, len(e.players)-1)
	for _, pid := range e.players {
		if pid == excludePid {
			continue
		}
		recipients = append(recipients, pid)
	}
	e.mu.RUnlock()

	for _, pid := range recipients {
		e.sendChannel <- sockt.Event[PID]{
			Type:    sockt.EventMessage,
			Subject: pid,
			Data:    []byte(fmt.Sprintf(`{"message": "Player %v %sed the room!"}`, excludePid, action)),
		}
	}
}
