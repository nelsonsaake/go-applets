package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"sync"
)

func index(writer http.ResponseWriter, request *http.Request) {
	t, _ := template.ParseFiles("./public/html/index.html")
	t.Execute(writer, "")
}

func GetPort() (port string) {
	port = ":8080"
	port = ":" + os.Getenv("PORT")
	return
}

func startServer(wg *sync.WaitGroup) {
	fmt.Println("note4u go server starting ...")
	defer wg.Done()

	server := http.Server{Addr: GetPort()}
	fmt.Println("Serving at port ", GetPort(), "...")

	setupDefaultRoute()
	setupUserRoutes()
	setupNoteRoutes()

	fmt.Println("starting server ...")
	server.ListenAndServe()
}

func startCmd(wg *sync.WaitGroup) {
	defer wg.Done()
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go startServer(&wg)
	go startCmd(&wg)
	wg.Wait()
}
