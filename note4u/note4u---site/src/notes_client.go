package n4u_admin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

//
func notesClientGet() (notes []Note, err error) {

	getUrl := g_url + "notes/"

	resp, err := http.Get(getUrl)

	if err != nil {
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return
	}

	err = json.Unmarshal(body, &notes)

	return
}

func notesClientDelete(id string) (body string, err error) {

	fmt.Println("requesting as client, deleted notes/" + id)

	client := &http.Client{}

	delUrl := g_url + "notes/" + id

	req, err := http.NewRequest("DELETE", delUrl, nil)

	if err != nil {
		return
	}

	resp, err := client.Do(req)

	if err != nil {
		return
	}

	bytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return
	}

	body = string(bytes)

	return
}

func notesClientPost(title, in_body, userId string) (body string, err error) {

	fmt.Println("requesting as client, create a new note, title=" + title)

	postUrl := g_url + "notes/"

	resp, err := http.PostForm(postUrl,
		url.Values{
			"title":  {title},
			"body":   {in_body},
			"userId": {userId},
		},
	)

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

func notesClientPut(id, title, in_body, userId string) (body string, err error) {

	fmt.Println("requesting as client, put notes/" + id)

	putUrl := g_url + "notes/" + id

	fmt.Println("requesting as client, update a note, title=" + title)

	resp, err := http.PostForm(putUrl,
		url.Values{
			"title":  {title},
			"body":   {in_body},
			"userId": {userId},
		},
	)

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
