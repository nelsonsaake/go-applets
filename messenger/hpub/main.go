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
	conn, _, _, err = ws.DefaultDialer.Dial(context.Background(), "ws://pubsubps.herokuapp.com/pub?id=foobar")
	if err != nil {
		// handle error
		err = fmt.Errorf("Error initiating socket handshake: %s", err.Error())
		panic(err)
	}

	fmt.Println("init...")
}

func writing(wg *sync.WaitGroup) {
	defer wg.Done()
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
	wg.Add(1)
	// go reading()
	go writing(wg)
	wg.Wait()
}
