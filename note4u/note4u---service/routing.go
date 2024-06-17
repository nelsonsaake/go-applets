package main

import (
	"fmt"
	"net/http"
	"projects/note4u---service/src"

	"github.com/rs/cors"
)

var c *cors.Cors

func init() {
	c = cors.AllowAll()
}

func setupDefaultRoute() {
	// default route
	http.Handle("/", c.Handler(http.HandlerFunc(index)))
	fmt.Println("Registered root path ...")
}

func setupUserRoutes() {
	// route dealing with users
	http.Handle("/login", c.Handler(http.HandlerFunc(src.Login)))
	fmt.Println("Registered /login path ...")

	http.Handle("/current-user", c.Handler(src.IsAuthorised(src.UsersGet)))
	fmt.Println("Registered /current-user path ...")

	http.Handle("/users", c.Handler(src.IsAuthorised(src.UsersHandler)))
	fmt.Println("Registered /users path ...")
}

func setupNoteRoutes() {
	// route dealing with notes
	http.Handle("/notes", c.Handler(src.IsAuthorised(src.NotesHandler)))
	fmt.Println("Registered /notes path ...")

	http.Handle("/notes/", c.Handler(src.IsAuthorised(src.NotesIdHandler)))
	fmt.Println("Registered /notes/ path ...")

	http.Handle("/notes/favourite", c.Handler(src.IsAuthorised(src.NotesUpdateFavorite)))
	fmt.Println("Registered /notes/favourite path ...")
}
