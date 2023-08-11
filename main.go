package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Just in case a wrong route enters this handler.
	// Might not be a practical use case, don't know yet; still learning.
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		log.Printf("Path not found: %v\n", r.URL.Path)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusMethodNotAllowed)
		log.Printf("Unsupported method: %v\n", r.Method)
		return
	}

	fmt.Fprintln(w, "Hello World!")
	log.Println("Fetched /hello")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Form was not able to be parsed properly\nError: %v", err)
		log.Println("Invalid form data received")
		return
	}

	if r.Method == "GET" {
		log.Println("Serving /form")
		http.ServeFile(w, r, "./static/form.html")
		return
	}

	if r.Method == "POST" {
		fmt.Fprintln(w, "POST request successful")
		name := r.FormValue("name")
		address := r.FormValue("address")
		fmt.Fprintf(w, "Name: %s\nAddress: %s\n", name, address)

		log.Println("Form data parsed successfully")
		return
	}
}

func main() {
	// specify the root directory for website
	fileserver := http.FileServer(http.Dir("./static"))

	// register functions (handlers) for a given pattern (http request)
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	// message to display in console when server starts.
	fmt.Println("Starting server at port http://0.0.0.0...")

	// ListenAndServe always returns a non-nil error
	// checks if port is busy
	// shut down server if port is already in use
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
