package main

import (
	"io"
	"net/http"
	"time"
)

func handleHelloWorld(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello Again World !")
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "The time is: "+time.Now().String())
}
