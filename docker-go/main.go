package main

import (
	"fmt"
	"net/http"
)

var counter int

func handler(w http.ResponseWriter, r *http.Request) {
	counter++
	fmt.Fprintf(w, "You are visitor number %d\n", counter)

	path := r.URL.Path
	if len(path) >= 2 {
		subPath := path[0:2]
		fmt.Fprintf(w, "First two characters of path: %s\n", subPath)
	} else {
		fmt.Fprintf(w, "Path is too short to extract first two characters.\n")
	}
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}
