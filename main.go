package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	fmt.Printf("Starting web server (0.0.0.0:8080)...")
	http.HandleFunc("/", handler)
	http.ListenAndServe("0.0.0.0:8080", nil)
}
