package httphandler

import (
	"io"
	"net/http"
	"time"
)

func TimeHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "The time is: "+time.Now().String())
}

func PingHanlder(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Pong")
}
