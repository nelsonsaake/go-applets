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

func initWS(ct ClientType) net.Conn {
	url := getWSURL(ct)
	ctx := context.Background()

	conn, _, _, err := ws.DefaultDialer.Dial(ctx, url)
	if err != nil {
		err = fmt.Errorf("Error initiating socket handshake: %s", err.Error())
		return nil
	}

	return conn
}

func pingpong(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Doing the ping pong thing...")
	for {
		msg, opCode, err := wsutil.ReadServerData(conn)
		if err != nil {
			// handle error
			fmt.Println("Error reading server data: ", err.Error())
			return
		}

		if !(opCode == ws.OpPing || opCode == ws.OpPong) {
			fmt.Print(string(msg))
		}

		conn.Write(ws.CompiledPing)
	}
}

func writing(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("writing...")

	for {
		ln, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println("Error reading message from standard input: ", err.Error())
			return
		}

		err = wsutil.WriteClientMessage(conn, ws.OpText, []byte(ln))
		if err != nil {
			fmt.Println("Error writing client message: ", err.Error())
			return
		}
	}
}

func main() {

	ct := getCT()

	conn := initWS(ct)
	if conn == nil {
		return
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go pingpong(conn, &wg)
	if ct == pub {
		go writing(conn, &wg)
	}

	wg.Wait()
}
