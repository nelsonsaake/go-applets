package n4u_admin

import (
	"fmt"
	"net/http"

	"github.com/rs/cors"
)

// url variable holds the location of the api, the api being managed here

var g_url string

// initialisation
// initialise the url

func init() {
	g_url = "https://localhost:8080/"
}

// the js, css, bootstrap... that will be used by the web pages delievered by the
// img handler, for the image directory, to serve image files for that directory
//
// css	for css files
// js	for javascript files
// bsjq for bootstrap files

func serveHtmlComponents() {
	fmt.Println("note4u manager registering html components ...")

	// the http.FileServer(http.Dir("./public/img/"))
	//
	// the http.FileServer returns the handler to the directory
	// the handler supplies the files in the directory when requested
	//
	// the "./public/img/", is the location of the directory
	//
	// the location of the file is relative to the native standalone file produced
	// when the application is built, it will produce a native file, a single standalone file

	img := http.FileServer(http.Dir("./public/img/"))
	http.Handle("/img/", http.StripPrefix("/img/", img))

	css := http.FileServer(http.Dir("./public/css/"))
	http.Handle("/css/", http.StripPrefix("/css/", css))

	js := http.FileServer(http.Dir("./public/js/"))
	http.Handle("/js/", http.StripPrefix("/js/", js))

	bsjq := http.FileServer(http.Dir("./public/bsjq/"))
	http.Handle("/bsjq/", http.StripPrefix("/bsjq/", bsjq))

	fmt.Println("note4u manager registering html components, success ...")
}

// add more routes
func RegisterRoutes() {
	fmt.Println("note4u manager routes registering ...")

	serveHtmlComponents()

	c := cors.AllowAll()

	http.Handle("/users/", c.Handler(http.HandlerFunc(users)))

	fmt.Println("note4u manager routes registering, success ...")
}
