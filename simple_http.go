package main

import (
	"log"
	"net/http"
)

func handle(w http.ResponseWriter, request *http.Request) {
	data := []byte(`{"Method Not Allowed"}`)
	if request.Method == "GET" {
		data := []byte(`{"name : MarkTBSS", "method : GET"}`)
		w.Write(data)
		return
	}
	if request.Method == "POST" {
		data = []byte(`{"name : MarkTBSS", "method : POST"}`)
		w.Write(data)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write(data)
}

func main() {
	http.HandleFunc("/", handle)
	log.Println("Starting HTTP server : 2565")
	log.Fatal(http.ListenAndServe(":2565", nil))
}
