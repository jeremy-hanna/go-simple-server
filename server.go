package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	mux := simpleMux()
	mux.HandleFunc("/hello", simpleHandler)
	// nil uses the DefaultServeMux
	log.Fatal(http.ListenAndServe(":7896", mux))
}

func simpleHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	io.WriteString(w, "Hello, world!\n")
}

func simpleMux() *http.ServeMux {
	m := http.NewServeMux()
	return m
}
