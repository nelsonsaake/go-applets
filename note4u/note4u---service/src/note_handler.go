package src

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strconv"
)

// admin
func notesAll(writer http.ResponseWriter, request *http.Request) {
	note := Note{}
	notes, err := note.All()
	if err != nil {
		errorMsg(writer, err)
		return
	}

	out, err := json.MarshalIndent(notes, "", "\t")
	if err != nil {
		errorMsg(writer, err)
		return
	}

	fmt.Fprintln(writer, string(out))
	writer.Header().Set("Content-Type", "application/json")
}

// admin
func notesGet(writer http.ResponseWriter, request *http.Request) {

	idstr := path.Base(request.URL.Path)
	id, err := strconv.Atoi(idstr)
	if err != nil {
		notesAll(writer, request)
		return
	}

	user, err := GetRequesttingUser(request)
	if err != nil {
		errorMsg(writer, err)
		return
	}

	note, err := Notes(id, user.Id)
	if err != nil {
		errorMsg(writer, err)
		return
	}

	out, _ := json.MarshalIndent(note, "", "\t")
	if err != nil {
		errorMsg(writer, err)
		return
	}

	fmt.Fprintln(writer, string(out))
	writer.Header().Set("Content-Type", "application/json")
}

// /notes post new note
func notesPost(writer http.ResponseWriter, request *http.Request) {

	request.ParseForm()

	user, err := GetRequesttingUser(request)
	if err != nil {
		errorMsg(writer, err)
		return
	}

	note := Note{
		Title: request.Form["title"][0],
		Body:  request.Form["body"][0],
		User:  &user,
	}

	err = note.Create()
	if err != nil {
		errorMsg(writer, err)
		return
	}

	writer.WriteHeader(200)
	return
}

// /notes/id put note
func notesPut(writer http.ResponseWriter, request *http.Request) {

	idstr := path.Base(request.URL.Path)
	id, err := strconv.Atoi(idstr)
	if err != nil {
		errorMsg(writer, err)
		return
	}

	user, err := GetRequesttingUser(request)
	if err != nil {
		errorMsg(writer, err)
		return
	}

	note, err := Notes(id, user.Id)
	if err != nil {
		errorMsg(writer, err)
		return
	}

	request.ParseForm()
	note.Title = request.Form["title"][0]
	note.Body = request.Form["body"][0]
	note.User = &user

	err = note.Update()
	if err != nil {
		errorMsg(writer, err)
		return
	}

	writer.WriteHeader(200)
	return
}

// /notes/id delete deletes note
func notesDelete(writer http.ResponseWriter, request *http.Request) {

	idstr := path.Base(request.URL.Path)
	id, err := strconv.Atoi(idstr)
	if err != nil {
		errorMsg(writer, err)
		return
	}

	user, err := GetRequesttingUser(request)
	if err != nil {
		errorMsg(writer, err)
		return
	}

	note, err := Notes(id, user.Id)
	if err != nil {
		errorMsg(writer, err)
		return
	}

	note.User = &user

	err = note.Delete()
	if err != nil {
		errorMsg(writer, err)
		return
	}

	writer.WriteHeader(200)
	return
}

// we are not expecting Id
// /notes search
// /notes post new note
func NotesHandler(writer http.ResponseWriter, request *http.Request) {

	fmt.Println("notes method ", request.Method)

	switch request.Method {

	case "GET":
		NotesSearch(writer, request)

	case "POST":
		notesPost(writer, request)

	default:
	}
}

// we are expecting an Id
// /notes/id put note
// /notes/id delete deletes note
func NotesIdHandler(writer http.ResponseWriter, request *http.Request) {

	fmt.Println("notes method ", request.Method)

	switch request.Method {

	case "PUT":
		notesPut(writer, request)

	case "DELETE":
		notesDelete(writer, request)

	default:
	}
}

// /notes search
func NotesSearch(writer http.ResponseWriter, request *http.Request) {
	user, err := GetRequesttingUser(request)
	if err != nil {
		errorMsg(writer, err)
		return
	}

	request.ParseForm()
	keyword := request.Form["search"][0]
	userId := user.Id
	notes, err := Search(keyword, userId)
	if err != nil {
		errorMsg(writer, err)
		return
	}

	out, err := json.MarshalIndent(notes, "", "\t")
	if err != nil {
		errorMsg(writer, err)
		return
	}

	fmt.Fprintln(writer, string(out))
	writer.Header().Set("Content-Type", "application/json")
}

// /notes/favourite/id
func NotesUpdateFavorite(writer http.ResponseWriter, request *http.Request) {
	idstr := path.Base(request.URL.Path)
	id, err := strconv.Atoi(idstr)
	if err != nil {
		errorMsg(writer, err)
		return
	}

	user, err := GetRequesttingUser(request)
	if err != nil {
		errorMsg(writer, err)
		return
	}

	note, err := Notes(id, user.Id)
	if err != nil {
		errorMsg(writer, err)
		return
	}

	note.User.Id = user.Id

	request.ParseForm()
	note.Favourite, err = strconv.ParseBool(request.Form["favourite"][0])
	if err != nil {
		err = fmt.Errorf("Error trying to update favourite: %s", err.Error())
		errorMsg(writer, err)
		return
	}

	err = note.UpdateFavourite()
	if err != nil {
		errorMsg(writer, err)
		return
	}

	writer.WriteHeader(200)
	return
}
