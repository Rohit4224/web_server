package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "Parseform() err: %v", err)
		return
	}
	fmt.Fprintf(w, "Post request Successful\n")
	name := r.FormValue("name")
	address := r.FormValue("42")
	fmt.Fprintf(w, "Name = %s\n 42 = %s\n", name, address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found, sorry", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method unsupported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello 42ers!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)   // handling form data submitted via a POST request to the "/form" path
	http.HandleFunc("/hello", helloHandler) // for handling a GET request to the "/hello" path

	fmt.Println("Starting your server at port 8080")
	err := http.ListenAndServe(":8080", nil) // the ListenAndServe function from the http package to start the server and listen for incoming requests on port 8080
	if err != nil {
		log.Fatal(err)
	}
}

// func main() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
// 	})

// 	//log.Println("Starting server on :8080")
// 	var fs = http.FileServer(http.Dir("static/"))
// 	http.Handle("/static/", http.StripPrefix("/static/", fs))
// 	var err = http.ListenAndServe(":80", nil)
// 	if err != nil {
// 		log.Fatal("Failed to start server: ", err)
// 	}
// }
