package main

import "net/http"

func main() {
	http.HandleFunc("/", handleHelloWorld)
	http.ListenAndServe(":8080", nil)
}
