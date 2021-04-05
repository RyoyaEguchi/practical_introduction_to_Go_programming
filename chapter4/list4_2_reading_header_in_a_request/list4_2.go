package main

import (
	"fmt"
	"net/http"
)

func headers(w http.ResponseWriter, r *http.Request) {
	h1 := r.Header
	h2 := r.Header["Accept-Encoding"]
	h3 := r.Header.Get("Accept-Encoding")
	fmt.Fprintln(w, h1, h2, h3)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/headers", headers)
	server.ListenAndServe()
}
