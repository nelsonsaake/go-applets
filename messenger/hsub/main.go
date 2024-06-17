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
	conn, _, _, err = ws.DefaultDialer.Dial(context.Background(), "ws://pubsubps.herokuapp.com/sub?id=foobar")
	if err != nil {
		// handle error
		err = fmt.Errorf("Error initiating socket handshake: %s", err.Error())
		panic(err)
	}

	fmt.Println("init...")
}

func writing() {
	fmt.Println("writing...")
	for {
		ln, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			// handle error
			fmt.Println("Error reading message from standard input: ", err.Error())
			return
		}

		// send message
		err = wsutil.WriteClientMessage(conn, ws.OpText, []byte(ln))
		if err != nil {
			// handle error
			fmt.Println("Error writing client message: ", err.Error())
			return
		}
	}
}

func reading(wg *sync.WaitGroup) {
	defer wg.Done()
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
	wg.Add(1)
	go reading(wg)
	// go writing()
	wg.Wait()
}
