package discordhandler

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var PLAYER_IDS = []string{"344869437750509568", "394094945985757185", "473510771796606996", "534942685610508298"}

func HandleCsAssemble(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content == "!cs" {
		message := "Playing CS ?"
		for _, player := range PLAYER_IDS {
			message += fmt.Sprintf("<@%s> ", player)
		}

		_, _ = s.ChannelMessageSend(m.ChannelID, message)
	}
}
