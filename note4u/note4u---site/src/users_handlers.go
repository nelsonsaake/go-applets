// this part of the application communicates with user of this application
// gets action the user wants to performed on the users table of the api and attempts to work on it
//
// in this file, two types of uses will be referenced,
// the users of the application and the users in the users table of the api we are managing
// users here will mean the users of this application
// and api-users will be users of the api we are trying to manage

package n4u_admin

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"path"
	"strconv"
)

// lrr is the last request response and it is stored as a string of characters
var lrr string

// initialises lrr
func init() {
	lrr = ""
}

// usersGet is called if the users requests for api-user information
// it is also called if the method to be performed on the api-users is not handled

// it gets data form usersClientGet
// create template using files found in the public directory
// it combines the template and the data to arrive at a page
// it returns the page

func usersGet(writer http.ResponseWriter, request *http.Request) {

	// write to console what is happening
	fmt.Println("users page requested ...")

	// the files to create the template
	tmpls := []string{"./public/html/layout.html", "./public/html/users.html"}

	// get the data from api-users
	users, err := usersClientGet()

	// 	if the was an error
	// 	write the error to console
	// 	write the error to the page
	// 	then exit
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(writer, err.Error())
		return
	}

	// creates a users table
	// this has all the information the template will require to build the page
	usersTable := UsersTable{Name: "Users", Users: users, LastRequestResponse: lrr}
	lrr = ""

	// create template
	t, err := template.ParseFiles(tmpls...)

	// if there was an error when creating a template
	// ... to console
	// ... to page
	// exit
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(writer, err.Error())
		return
	}

	// combine the data and the template to produce a page
	//
	t.ExecuteTemplate(writer, "layout", usersTable)
}

//
//
func usersDelete(writer http.ResponseWriter, request *http.Request) {

	fmt.Println("users delete requested")

	idstr := path.Base(request.URL.Path)
	_, err := strconv.Atoi(idstr)

	if err != nil {
		err = errors.New("id provided is not an integer")
		fmt.Println(err)
		fmt.Fprintf(writer, err.Error())
		return
	}

	fmt.Println("id =", idstr)

	lrr, err = usersClientDelete(idstr)

	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(writer, err.Error())
		return
	}
}

func usersPost(writer http.ResponseWriter, request *http.Request) {

	fmt.Println("requesting the creation of a new user")

	// first get values
	email := request.FormValue("email")
	pwd := request.FormValue("password")

	// use the values to compose request to send
	var err error
	lrr, err = usersClientPost(email, pwd)

	// just exit, the json will redirect back
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(writer, err.Error())
		return
	}
}

func usersPut(writer http.ResponseWriter, request *http.Request) {

	// first get id

	// get values

	// compose request

	// send request

	// exit
}

func users(writer http.ResponseWriter, request *http.Request) {

	method := request.Method

	fmt.Println("users ", method)

	switch method {

	case "GET":
		usersGet(writer, request)

	case "POST":
		usersPost(writer, request)

	case "PUT":
		usersPut(writer, request)

	case "DELETE":
		usersDelete(writer, request)

	default:
		usersGet(writer, request)
	}
}
