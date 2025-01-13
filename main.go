package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ShervilG/spinboard-go/httphandler"
	"github.com/bwmarrin/discordgo"
)

var discordSession *discordgo.Session

func main() {
	discordBotToken := os.Getenv("BUNTY_BOT_TOKEN")
	discordSession, err := discordgo.New("Bot " + discordBotToken)
	if err != nil {
		fmt.Printf("Error creating Discord session: %v\n", err)
	}

	discordSession.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMessages | discordgo.IntentsDirectMessages | discordgo.IntentsGuildPresences
	discordSession.Open()
	defer discordSession.Close()

	http.HandleFunc("/", httphandler.HandleHelloWorld)
	http.HandleFunc("/time", httphandler.TimeHandler)

	http.ListenAndServe(":8080", nil)
}
