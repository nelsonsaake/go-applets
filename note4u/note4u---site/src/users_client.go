// the note4u manager is used to maintain/manage/check the note4u api
//
// this part will create, retrieve, update and delete from the users table of the api

package n4u_admin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// send a request the api we are trying to manage
// send the response
// the response will be in json
// convert the response into some slice of users
// combine the users data with the template
// and return a html page
func usersClientGet() (users []User, err error) {

	// we want data from the users table
	// the users table is accessible at the /users/
	getUrl := g_url + "users/"

	// we send request
	resp, err := http.Get(getUrl)

	// if we get any error we exit
	if err != nil {
		return
	}

	// make sure to close the response body, else that will create a leak
	defer resp.Body.Close()

	// extract the response body
	body, err := ioutil.ReadAll(resp.Body)

	// if there was an error in getting the body we exit
	if err != nil {
		return
	}

	// we parse the json body to get the users data
	err = json.Unmarshal(body, &users)

	// we don't check for error because we are returning anyways
	// and we are expecting the caller to check for error when this function returns

	// when we exit we go back to the function that called this one
	return
}

// this function attempts to delete a particular user from the api
func usersClientDelete(id string) (body string, err error) {

	// write to the console what is happening
	fmt.Println("requesting as client, deleted users/" + id)

	// create a client
	// client is that struct that will help us send the request
	client := &http.Client{}

	// modifier the url to point to that particular user
	delUrl := g_url + "users/" + id

	// create a request
	// create will hold informations like :
	//  the method/action we would want to perform on the resource,"DELETE"
	//  the location of the resource, delUrl
	req, err := http.NewRequest("DELETE", delUrl, nil)

	// if there was an error making the request: we exit
	if err != nil {
		return
	}

	// now we send request and get the response
	resp, err := client.Do(req)

	// if there was an error ...
	if err != nil {
		return
	}

	// we get the response
	bytes, err := ioutil.ReadAll(resp.Body)

	// ...
	if err != nil {
		return
	}

	// convert the response to the type string
	// the response type is by default bytes
	// bytes is different from string basically because of the operation that can be performed onn them
	body = string(bytes)

	// ...///
	return
}

// ?
func usersClientPost(email string, pwd string) (body string, err error) {

	fmt.Println("requesting as client, create a user, email=" + email)

	postUrl := g_url + "users/"

	resp, err := http.PostForm(postUrl, url.Values{"key": {"Value"}, "email": {email}, "password": {pwd}})

	if err != nil {
		return
	}

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return
	}

	body = string(bytes)

	return
}

// ?
func usersClientPut(id string, email string, pwd string) (body string, err error) {

	fmt.Println("requesting as client, put users/" + id)

	putUrl := g_url + "users/" + id

	fmt.Println("requesting as client, update a new user, email=" + email)

	resp, err := http.PostForm(putUrl, url.Values{"key": {"Value"}, "email": {email}, "password": {pwd}})

	if err != nil {
		return
	}

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return
	}

	body = string(bytes)

	return
}
