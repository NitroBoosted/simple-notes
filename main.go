package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	h1 := func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello World")
	}
	http.HandleFunc("/", h1)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
