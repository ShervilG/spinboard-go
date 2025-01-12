package main

import "net/http"

func main() {
	http.HandleFunc("/", handleHelloWorld)
	http.HandleFunc("/time", timeHandler)

	http.ListenAndServe(":8080", nil)
}
