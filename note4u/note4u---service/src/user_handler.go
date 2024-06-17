package src

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strconv"
)

// /current-user get
func UsersGet(writer http.ResponseWriter, request *http.Request) {

	user, err := GetRequesttingUser(request)
	if err != nil {
		errorMsg(writer, err)
		return
	}

	out, err := json.MarshalIndent(user, "", "\t")
	if err != nil {
		errorMsg(writer, err)
		return
	}

	fmt.Fprintln(writer, string(out))
	writer.Header().Set("Content-Type", "application/json")
}

// /users post
func usersPost(writer http.ResponseWriter, request *http.Request) {

	request.ParseForm()

	user := User{
		Email:    request.Form["email"][0],
		Password: request.Form["password"][0],
	}
	err := user.Create()
	if err != nil {
		errorMsg(writer, err)
		return
	}

	writer.WriteHeader(200)
	return
}

// admin
func usersPut(writer http.ResponseWriter, request *http.Request) {

	idstr := path.Base(request.URL.Path)
	id, err := strconv.Atoi(idstr)
	if err != nil {
		errorMsg(writer, err)
		return
	}

	user, err := Users(id)
	if err != nil {
		errorMsg(writer, err)
		return
	}

	request.ParseForm()
	user.Email = request.Form["email"][0]
	user.Password = request.Form["password"][0]

	err = user.Update()
	if err != nil {
		errorMsg(writer, err)
		return
	}

	writer.WriteHeader(200)
	return
}

// admin
func usersDelete(writer http.ResponseWriter, request *http.Request) {

	idstr := path.Base(request.URL.Path)
	id, err := strconv.Atoi(idstr)
	if err != nil {
		errorMsg(writer, err)
		return
	}

	user, err := Users(id)
	if err != nil {
		errorMsg(writer, err)
		return
	}

	err = user.Delete()
	if err != nil {
		errorMsg(writer, err)
		return
	}

	writer.WriteHeader(200)
	return
}

//
// /users post
func UsersHandler(writer http.ResponseWriter, request *http.Request) {

	fmt.Println("users method ", request.Method)

	switch request.Method {

	case "POST":
		usersPost(writer, request)

	default:
	}
}

// /login post
func Login(writer http.ResponseWriter, request *http.Request) {

	request.ParseForm()
	var email string
	var password string

	if request.Form["email"] != nil {
		email = request.Form["email"][0]
	} else {
		fmt.Fprintf(writer, "Error logining in, no email provided")
		return
	}

	if request.Form["password"] != nil {
		password = request.Form["password"][0]
	} else {
		fmt.Fprintf(writer, "Error logining in, no password provided")
		return
	}

	user, err := GetUser(email, password)

	if err != nil {
		fmt.Fprintf(writer, "Error getting user: %s", err.Error())
		return
	}

	tokenString, err := GenerateJWT(user)

	if err != nil {
		fmt.Fprintf(writer, "Error logging: %s", err.Error())
		return
	}

	writer.WriteHeader(200)
	writer.Header().Set("Token", tokenString)
}
