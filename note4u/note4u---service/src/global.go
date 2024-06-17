package src

import (
	"fmt"
	"net/http"
)

func errorMsg(writer http.ResponseWriter, err error) {
	fmt.Fprintf(writer, err.Error())
	fmt.Println(err)
}
