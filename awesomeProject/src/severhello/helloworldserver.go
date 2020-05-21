package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "<h1>hello your %s!</h>",
			request.FormValue("name"))
	})

	http.ListenAndServe(":5555", nil)
}
