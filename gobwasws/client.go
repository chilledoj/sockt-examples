package gobwasws

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"log"
	"time"
)

func RunGobwasWsClient(ctx context.Context) error {
	c, _, _, err := ws.DefaultDialer.Dial(ctx, "ws://localhost:9090/api/ws")
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	fmt.Println("Connected!")
	if _, err := c.Write([]byte(`{"message": "hi"}`)); err != nil {
		return err
	}
	go func(ctx context.Context) {
		for {
			fmt.Println("Reading...")

			message, _, err := wsutil.ReadServerData(c)
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
