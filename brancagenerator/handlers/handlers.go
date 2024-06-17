package handlers

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/gin-gonic/gin/render"
	"github.com/hako/branca"
)

func Test(w http.ResponseWriter, r *http.Request) {

	tmpl := []string{
		"./static/html/testing.html",
	}

	t, err := template.ParseFiles(tmpl...)
	if err != nil {
		fmt.Println("Error parsing template file...")
	}
	t.Execute(w, "")
}

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl := []string{
		"./static/html/wrapper.html",
		"./static/html/header.html",
		"./static/html/content/index.html",
		"./static/html/footer.html",
	}

	t, err := template.ParseFiles(tmpl...)
	if err != nil {
		fmt.Println("Error parsing template file...")
	}
	t.ExecuteTemplate(w, "wrapper", "")
}

func Error(w http.ResponseWriter, r *http.Request) {
	tmpl := []string{
		"./static/html/wrapper.html",
		"./static/html/header.html",
		"./static/html/content/error.html",
		"./static/html/footer.html",
	}

	t, err := template.ParseFiles(tmpl...)
	if err != nil {
		fmt.Println("Error parsing template file...")
	}
	t.ExecuteTemplate(w, "wrapper", "")
}

func Help(w http.ResponseWriter, r *http.Request) {
	tmpl := []string{
		"./static/html/wrapper.html",
		"./static/html/header.html",
		"./static/html/content/help.html",
		"./static/html/footer.html",
	}

	t, err := template.ParseFiles(tmpl...)
	if err != nil {
		fmt.Println("Error parsing template file...")
	}
	t.ExecuteTemplate(w, "wrapper", "")
}

func About(w http.ResponseWriter, r *http.Request) {
	tmpl := []string{
		"./static/html/wrapper.html",
		"./static/html/header.html",
		"./static/html/content/about.html",
		"./static/html/footer.html",
	}

	t, err := template.ParseFiles(tmpl...)
	if err != nil {
		fmt.Println("Error parsing template file...")
	}
	t.ExecuteTemplate(w, "wrapper", "")
}

type BrancaAction int

const (
	Encode BrancaAction = 0
	Decode              = 1
)

type BrancaProcessorRequest struct {
	Action BrancaAction
	Key    string
	Text   string
}

func BrancaProcessor(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		//

		bpReq := BrancaProcessorRequest{}
		brc := branca.NewBranca(bpReq.Key)
		res := ""
		var err error

		switch bpReq.Action {
		case Encode:
			res, err = brc.EncodeToString(bqReq.Text)
		case Decode:
			res, err = brc.DecodeToString(bqReq.Text)
		}

		if err != nil {

			errMsg := fmt.Sprint("Error processing branca ", err.Error())
			fmt.Println(errMsg)
			formatter.Json(w, http.StatusBadRequest, struct{ Error string }{errMsg})
			return
		}
	}
}
