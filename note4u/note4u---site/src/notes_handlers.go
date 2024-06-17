package n4u_admin

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"path"
	"strconv"
)

func notesGet(writer http.ResponseWriter, request *http.Request) {

	fmt.Println("notes page requested ...")

	tmpls := []string{"./public/html/layout.html", "./public/html/notes.html"}

	notes, err := notesClientGet()

	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(writer, err.Error())
		return
	}

	notesTable := NotesTable{Name: "Notes", Notes: notes, LastRequestResponse: lrr}
	lrr = ""

	t, err := template.ParseFiles(tmpls...)

	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(writer, err.Error())
		return
	}

	t.ExecuteTemplate(writer, "layout", notesTable)
}

func notesDelete(writer http.ResponseWriter, request *http.Request) {

	fmt.Println("notes delete requested")

	idstr := path.Base(request.URL.Path)
	_, err := strconv.Atoi(idstr)

	if err != nil {
		err = errors.New("id provided is not an integer")
		fmt.Println(err)
		fmt.Fprintf(writer, err.Error())
		return
	}

	fmt.Println("id =", idstr)

	lrr, err = notesClientDelete(idstr)

	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(writer, err.Error())
		return
	}
}

func notesPost(writer http.ResponseWriter, request *http.Request) {

	fmt.Println("requesting the creation of a new note")

	title := request.FormValue("title")
	body := request.FormValue("body")
	userId := request.FormValue("userId")

	var err error
	lrr, err = notesClientPost(title, body, userId)

	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(writer, err.Error())
		return
	}
}

func notesPut(writer http.ResponseWriter, request *http.Request) {

	// first get id

	// get values

	// compose request

	// send request

	// exit
}

func notes(writer http.ResponseWriter, request *http.Request) {

	method := request.Method

	fmt.Println("notes ", method)

	switch method {

	case "GET":
		notesGet(writer, request)

	case "POST":
		notesPost(writer, request)

	case "PUT":
		notesPut(writer, request)

	case "DELETE":
		notesDelete(writer, request)

	default:
		notesGet(writer, request)
	}
}
