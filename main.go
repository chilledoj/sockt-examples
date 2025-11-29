package main

import (
	"context"
	"fmt"
	"github.com/chilledoj/sockt"
	"github.com/chilledoj/sockt-examples/coderws"
	"github.com/chilledoj/sockt-examples/engine"
	"github.com/chilledoj/sockt-examples/tcpserve"
	"github.com/chilledoj/sockt-examples/types"
	"log"
	"runtime"
	"time"

	"os"
	"os/signal"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide server type as first argument")
	}
	serverType := os.Args[1]
	if serverType == "" {
		serverType = "coderws"
	}

	err := run(serverType)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Second * 3)

	log.Printf("GO ROUTINES: %d", runtime.NumGoroutine())

	os.Exit(0)
}

func run(serverType string) error {
	log.Printf("Starting server of type: %s", serverType)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	var roomID = "room1"

	engine := engine.NewEngine[types.PlayerID](ctx)

	room := sockt.NewRoom[types.RoomID, types.PlayerID](ctx, roomID, engine)
	go room.Run()

	var srv server

	srv, err := serverFactory(serverType, room)
	if err != nil {
		return err
	}

	go func() {
		log.Println("Starting server...")
		err := srv.Start()
		if err != nil {
			log.Println(err)
		}
	}()

	<-ctx.Done()
	log.Println("Shutting down...")
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), time.Second*5)
	defer shutdownCancel()
	srv.Stop(shutdownCtx)
	room.Stop()
	engine.Stop()

	return ctx.Err()
}

type server interface {
	Start() error
	Stop(ctx context.Context) error
}

func serverFactory(serverType string, room *sockt.Room[types.RoomID, types.PlayerID]) (server, error) {
	switch serverType {
	case "coderws":
		return coderws.NewCoderServer(room), nil
	case "tcp":
		return tcpserve.NewTcpServer(room)
	case "":
		return nil, fmt.Errorf("no server type provided")
	default:
		return nil, fmt.Errorf("unknown server type: %s", serverType)
	}
}
