package coderws

import (
	"context"
	"fmt"
	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
	"time"
)

func RunCoderClient(ctx context.Context) error {
	c, _, err := websocket.Dial(ctx, "ws://localhost:9090/api/ws", nil)
	if err != nil {
		return err
	}
	defer c.CloseNow()
	fmt.Println("Connected!")

	err = wsjson.Write(ctx, c, `{"message": "hi"}`)
	if err != nil {
		return err
	}
	time.Sleep(time.Second * 2)

	go func(ctx context.Context) {
		for {
			fmt.Println("Reading...")
			var v any
			err := wsjson.Read(ctx, c, &v)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("Received: ", v)
			if ctx.Err() != nil {
				fmt.Println("Exiting read loop...")
				return
			}
		}
	}(ctx)

	time.Sleep(time.Second * 5)

	fmt.Println("Closing connection...")

	return c.Close(websocket.StatusNormalClosure, "client closed connection")
}
