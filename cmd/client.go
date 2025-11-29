package main

import (
	"context"
	"fmt"
	"github.com/chilledoj/sockt-examples/coderws"
	"github.com/chilledoj/sockt-examples/tcpserve"
	"log"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide server type as first argument")
	}
	serverType := os.Args[1]
	if serverType == "" {
		serverType = "coderws"
	}

	if err := runClient(serverType); err != nil {
		fmt.Println(err)
	}
}

func runClient(serverType string) error {
	log.Printf("Starting client of type: %s", serverType)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	switch serverType {
	case "coderws":
		return coderws.RunCoderClient(ctx)
	case "tcp":
		return tcpserve.RunTcpClient(ctx)
	case "":
		return fmt.Errorf("no server type provided")
	default:
		return fmt.Errorf("unknown server type: %s", serverType)
	}
	return nil
}
