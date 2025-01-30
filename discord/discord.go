package discord

import (
	"os"

	"github.com/bwmarrin/discordgo"
)

var ds *discordgo.Session

func InitSession() {
	discordBotToken := os.Getenv("BUNTY_BOT_TOKEN")
	ds, _ = discordgo.New("Bot " + discordBotToken)

	ds.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMessages | discordgo.IntentsDirectMessages | discordgo.IntentsGuildPresences | discordgo.IntentsGuildVoiceStates
}

func GetDiscordSession() *discordgo.Session {
	return ds
}
