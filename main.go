package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	//log.Println("Starting server on :8080")
	var fs = http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	var err = http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
