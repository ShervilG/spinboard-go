package main

import (
	"io"
	"net/http"
)

func handleHelloWorld(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, World!")
}
