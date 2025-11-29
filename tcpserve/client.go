package tcpserve

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/chilledoj/sockt-examples/types"
	"net"
	"os"
	"time"
)

func RunTcpClient(ctx context.Context) error {
	// Resolve the string address to a TCP address
	tcpAddr, err := net.ResolveTCPAddr("tcp4", ":9191")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Connect to the address with tcp
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return err
	}
	fmt.Println("Connected!")

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.Write([]byte(`{"message": "hi"}`))
	if err != nil {
		return err
	}
	time.Sleep(time.Second * 2)

	go func(ctx context.Context) {
		for {
			fmt.Println("Reading...")
			buf := make([]byte, 4096)
			n, err := conn.Read(buf)
			if err != nil {
				fmt.Println(err)
				return
			}
			var m types.Msg
			if err := json.Unmarshal(buf[:n], &m); err != nil {
				fmt.Printf("Error unmarshalling message: %v\n", err)
			}

			fmt.Println("Received: ", m)
			if ctx.Err() != nil {
				fmt.Println("Exiting read loop...")
				return
			}
		}
	}(ctx)

	time.Sleep(time.Second * 5)

	fmt.Println("Closing connection...")

	return conn.Close()
}
