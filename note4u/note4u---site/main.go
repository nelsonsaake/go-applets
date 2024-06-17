// note4u admin
// an interface to manage the note4u api

package main

import (
	"fmt"
	"net/http"
	n4u_admin "projects/note4u---site/src"

	"github.com/rs/cors"
)

// welcome page
func index(writer http.ResponseWriter, request *http.Request) {
	
	fmt.Println("requesting index/welcome page ...")
	fmt.Fprintf(writer, "note4u manager")
}

func main() {
	fmt.Println("note4u manager initializing ...")

	port := ":8081"
	server := http.Server{Addr: port}

	fmt.Println("note4u manager listening at port ", port, " ...")
	fmt.Println("note4u manager serving at port ", port, " ...")

	c := cors.AllowAll()

	n4u_admin.RegisterRoutes()

	http.Handle("/", c.Handler(http.HandlerFunc(index)))

	fmt.Println("note4u manager initialized ...")
	fmt.Println("note4u manager starting ...")

	server.ListenAndServe()
	fmt.Println("note4u manager stoped ...")
}

// the printing done may be redundunt but it's target at observability
// the printing is to say what is happening as specific as possible, more like loging
