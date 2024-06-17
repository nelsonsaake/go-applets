package main

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/rs/cors"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

var (
	ctx    context.Context
	err    error
	conn   net.Conn
	cancel context.CancelFunc
	pubURL string = "ws://pubsubps.herokuapp.com/pub?id=foobar"
	subURL string = "ws://pubsubps.herokuapp.com/sub?id=foobar"
)

//
func getWSURL(ct ClientType) string {
	switch ct {
	case pub:
		return pubURL
	case sub:
		return subURL
	default:
		fmt.Println("Error selecting url, bad client type.")
	}
	return ""
}

func initiateWS(ct ClientType) {
	url := getWSURL(ct)
	conn, _, _, err = ws.DefaultDialer.Dial(context.Background(), url)
	if err != nil {
		// handle error
		err = fmt.Errorf("Error initiating socket handshake: %s", err.Error())
		panic(err)
	}

	fmt.Println("init...")
}

//
func writing(wg *sync.WaitGroup) {
	defer wg.Done()
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

func writing2(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
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

//
func reading(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("reading...")

	for {
		msg, opCode, err := wsutil.ReadServerData(conn)
		if err != nil {
			// handle error
			fmt.Println("Error reading reading server data: ", err.Error())
			return
		}

		if !(opCode == ws.OpPing || opCode == ws.OpPong) {
			fmt.Println(string(msg))
		}
	}
}

func reading2(w http.ResponseWriter, conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("reading...")

	for {
		msg, opCode, err := wsutil.ReadServerData(conn)
		if err != nil {
			// handle error
			fmt.Println("Error reading reading server data: ", err.Error())
			return
		}

		if !(opCode == ws.OpPing || opCode == ws.OpPong) {
			fmt.Println(string(msg))
			fmt.Fprint(w, string(msg))
		}
	}
}

// do the ping pong thing AKA heart beat
func pingpong(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Doing the ping pong thing...")
	for {
		msg, opCode, err := wsutil.ReadServerData(conn)
		if err != nil {
			// handle error
			fmt.Println("Error reading reading server data: ", err.Error())
			return
		}

		if !(opCode == ws.OpPing || opCode == ws.OpPong) {
			fmt.Println(string(msg))
		}

		conn.Write(ws.CompiledPing)
	}
}

func pingpong2(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Doing the ping pong thing...")
	for {
		msg, opCode, err := wsutil.ReadServerData(conn)
		if err != nil {
			// handle error
			fmt.Println("Error reading reading server data: ", err.Error())
			return
		}

		if !(opCode == ws.OpPing || opCode == ws.OpPong) {
			fmt.Println(string(msg))
		}

		conn.Write(ws.CompiledPing)
	}
}

//
func getClientType() ClientType {
	ln, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		// handle error
		fmt.Println("Error reading message from standard input: ", err.Error())
		return bad
	}

	ln = strings.ToLower(ln)
	if ln == "pub" {
		return pub
	} else if ln == "sub" {
		return sub
	}

	return bad
}

func startClientActions(wg *sync.WaitGroup) {
	ct := getClientType()

	initiateWS(ct)
	go pingpong(wg)

	switch ct {
	case pub:
		go reading(wg)
	case sub:
		go writing(wg)
	default:
		fmt.Println("Error : bad client type.")
	}
	return
}

//
func hpub(w http.ResponseWriter, r *http.Request) {
	url := "ws://pubsubps.herokuapp.com/pub?id=foobar"
	conn, _, _, err = ws.DefaultDialer.Dial(context.Background(), url)
	if err != nil {
		err = fmt.Errorf("Error initiating socket handshake: %s", err.Error())
		return
	}

	var wg *sync.WaitGroup
	wg.Add(2)

	go pingpong2(conn, wg)
	go writing2(conn, wg)
}

func hsub(w http.ResponseWriter, r *http.Request) {
	url := "ws://pubsubps.herokuapp.com/sub?id=foobar"
	conn, _, _, err = ws.DefaultDialer.Dial(context.Background(), url)
	if err != nil {
		err = fmt.Errorf("Error initiating socket handshake: %s", err.Error())
		return
	}

	var wg *sync.WaitGroup
	wg.Add(2)

	go pingpong2(conn, wg)
	go reading2(w, conn, wg)
}

//
func defpub() {
	url := "ws://pubsubps.herokuapp.com/pub?id=foobar"
	conn, _, _, err = ws.DefaultDialer.Dial(context.Background(), url)
	if err != nil {
		err = fmt.Errorf("Error initiating socket handshake: %s", err.Error())
		return
	}

	var wg *sync.WaitGroup
	wg.Add(2)

	go pingpong2(conn, wg)
	go writing2(conn, wg)
}

// procedure
func defaultProcedure() {
	var wg *sync.WaitGroup
	wg.Add(2)
	startClientActions(wg)
	wg.Wait()
}

func server() {
	server := http.Server{
		Addr: ":8080",
	}
	c := cors.AllowAll()

	// http.Handle("/pub", c.Handler(http.HandlerFunc(hpub)))
	http.Handle("/sub", c.Handler(http.HandlerFunc(hsub)))

	server.ListenAndServe()
}

//
func main() {
	server()
}
