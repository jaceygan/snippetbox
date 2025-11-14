package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	port := flag.String("port", "4000", "Port to run the web server on")
	flag.Parse()

	mux := http.NewServeMux() // Always declare this. Never use the default mux.
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	log.Printf("Starting server on port %s", *port)
	err := http.ListenAndServe(":"+*port, mux)
	log.Fatal(err)
}
