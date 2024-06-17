package main

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"os"
	"sync"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

var (
	ctx    context.Context
	cancel context.CancelFunc
	conn   net.Conn
)

func init() {
	var err error
	conn, _, _, err = ws.DefaultDialer.Dial(context.Background(), "ws://localhost:8084")
	if err != nil {
		// handle error
		fmt.Println("Error initiating socket handshake: ", err.Error())
	}

	fmt.Println("init...")
}

func writing() {
	fmt.Println("writing...")
	for {
		ln, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		// send message
		err := wsutil.WriteClientMessage(conn, ws.OpText, []byte(ln))
		if err != nil {
			// handle error
			fmt.Println("Error writing client message: ", err.Error())
			return
		}
	}
}

func reading() {
	fmt.Println("reading...")
	for {
		msg, _, err := wsutil.ReadServerData(conn)
		if err != nil {
			// handle error
			fmt.Println("Error reading reading server data: ", err.Error())
			return
		}
		fmt.Println(string(msg))
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	// go reading()
	go writing()
	wg.Wait()
}
