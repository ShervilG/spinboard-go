package httphandler

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func HandleHelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request for hello world !")
	io.WriteString(w, "Hello Again World !")
}

func TimeHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "The time is: "+time.Now().String())
}
