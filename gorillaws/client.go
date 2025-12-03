package gorillaws

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

func RunGorillaClient(ctx context.Context) error {
	c, _, err := websocket.DefaultDialer.Dial("ws://localhost:9090/api/ws", nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	fmt.Println("Connected!")
	if err := c.WriteMessage(websocket.TextMessage, []byte(`{"message": "hi"}`)); err != nil {
		return err
	}
	go func(ctx context.Context) {
		for {
			fmt.Println("Reading...")
			_, message, err := c.ReadMessage()
			if err != nil {
				fmt.Println(err)
				return
			}
			var v any
			json.Unmarshal(message, &v) // ignore error here
			fmt.Println("Received: ", v)
			if ctx.Err() != nil {
				fmt.Println("Exiting read loop...")
				return
			}
		}
	}(ctx)

	time.Sleep(time.Second * 5)

	fmt.Println("Closing connection...")

	return c.Close()
}
