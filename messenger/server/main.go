package main

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/rs/cors"
)

var (
	ctx    context.Context
	cancel context.CancelFunc
	conn   net.Conn
)

func reading() {
	for {
		msg, op, err := wsutil.ReadClientData(conn)
		if err != nil {
			// handle error
			fmt.Println("Error: reading client data: ", err.Error())
			return
		}
		fmt.Println(string(msg))
		err = wsutil.WriteServerMessage(conn, op, []byte("Message Recieved"))
		if err != nil {
			// handle error
			fmt.Println("Error: writing server message: ", err.Error())
			return
		}
	}
}

func writing() {
	for {
		ln, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		err := wsutil.WriteServerMessage(conn, ws.OpText, []byte(ln))
		if err != nil {
			// handle error
			fmt.Println("Error: writing server message: ", err.Error())
			return
		}
	}
}

func upgrade(w http.ResponseWriter, r *http.Request) {
	var err error
	conn, _, _, err = ws.UpgradeHTTP(r, w)
	if err != nil {
		fmt.Println("Error: failed to upgrade connection", err.Error())
	}
	go reading()
	go writing()
}

func main() {
	server := http.Server{
		Addr: ":8084",
	}

	c := cors.AllowAll()
	http.Handle("/", c.Handler(http.HandlerFunc(upgrade)))

	server.ListenAndServe()
}
