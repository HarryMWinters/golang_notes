package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> Hello Index </h1>")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1> Hello About </h1>")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	fmt.Println("Program starting...")
	http.ListenAndServe(":3000", nil)
}
