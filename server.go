package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	mux := simpleMux()
	mux.HandleFunc("/hello", simpleHandler)
	mux.HandleFunc("/post", simpleQueryParam)
	// a pattern ending in a slash defines a subtree
	mux.HandleFunc("/post/", simplePathParam)
	// nil uses the DefaultServeMux
	log.Fatal(http.ListenAndServe(":7896", mux))
}

func simpleHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	io.WriteString(w, "Hello, world!\n")
}

func simpleQueryParam(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	p := req.URL.Query().Get("id")
	fmt.Fprintf(w, "Hello from id = %s", p)
}

func simplePathParam(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	p := req.URL.Path[len("/post/"):]
	fmt.Fprintf(w, "Hello from %s", p)
}

func simpleMux() *http.ServeMux {
	m := http.NewServeMux()
	return m
}
