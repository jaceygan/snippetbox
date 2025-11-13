package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	log.Print("Handling request for home page")
	w.Header().Add("Server", "Go")
	w.Header().Add("Server", "Go2")
	// .Add() multiple values for same header key
	// use .Set() to overwrite existing value instead
	w.Write([]byte("Hello from Snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Displaying a specific snippet with ID %d...", id)
	/* Alternative way:
	msg := fmt.Sprintf("Displaying snippet with ID %d", id)
	w.Write([]byte(msg))
	*/
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display form for creating a new snippet..."))
}

func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	// always update response header map before WriteHeader() or Write()
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Save a new snippet..."))
}
