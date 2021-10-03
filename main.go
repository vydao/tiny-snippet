package main

import (
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello World!")
	w.Write([]byte("Hello World!"))
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet"))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/create", createSnippet)
	mux.HandleFunc("/snippet", showSnippet)

	http.ListenAndServe(":8080", mux)
}
