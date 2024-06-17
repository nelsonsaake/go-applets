package service

import (
	"net/http"
	"projects/applets/brancagenerator/handlers"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func NewServer() *negroni.Negroni {

	n := negroni.Classic()

	mx := mux.NewRouter()

	initRoutes(mx)

	n.UseHandler(mx)

	return n
}

func initRoutes(mx *mux.Router) {

	// one route for the index page
	mx.HandleFunc("/", handlers.Index)

	// one route for the about page
	mx.HandleFunc("/about", handlers.About)

	// one route for the help page
	mx.HandleFunc("/help", handlers.Help)

	// one route for error page
	mx.NotFoundHandler = http.HandlerFunc(handlers.Error)

	// one route for testing page
	mx.HandleFunc("/test", handlers.Test)

	// image dir server
	imgHandler := http.FileServer(http.Dir("./static/img/"))
	mx.PathPrefix("/img/").Handler(http.StripPrefix("/img/", imgHandler))

}
