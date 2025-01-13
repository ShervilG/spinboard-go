package main

import (
	"net/http"
	"os"

	"github.com/ShervilG/spinboard-go/httphandler"
	"github.com/bwmarrin/discordgo"
)

var discordSession *discordgo.Session

func main() {
	discordBotToken := os.Getenv("BUNTY_BOT_TOKEN")
	discordSession, _ = discordgo.New(discordBotToken)
	defer discordSession.Close()

	http.HandleFunc("/", httphandler.HandleHelloWorld)
	http.HandleFunc("/time", httphandler.TimeHandler)

	http.ListenAndServe(":8080", nil)
}
