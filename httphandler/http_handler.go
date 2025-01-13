package httphandler

import (
	"io"
	"net/http"
	"time"
)

func HandleHelloWorld(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello Again World !")
}

func TimeHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "The time is: "+time.Now().String())
}
