package discordhandler

import "github.com/bwmarrin/discordgo"

func HandleHello(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!hello" {
		s.ChannelMessageSend(m.ChannelID, "Hello, "+m.Author.Username)
	}
}
