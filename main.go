package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlerFn)
	fmt.Println("Listening on :8080")
	_ = http.ListenAndServe(":8080", nil)
}

func handlerFn(w http.ResponseWriter, r *http.Request) {
	part := r.URL.Path[1:]

	if len(part) == 0 {
		part = "World"
	}

	_, _ = fmt.Fprintf(w, "Hello, %s!\n", part)
}