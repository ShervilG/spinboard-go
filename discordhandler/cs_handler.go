package discordhandler

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var PLAYER_IDS = []string{"344869437750509568", "394094945985757185", "473510771796606996", "534942685610508298", "694631299323002890"}

func HandleCsAssemble(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!cs" {
		message := "Playing CS ?"
		for _, player := range PLAYER_IDS {
			if player != m.Author.ID {
				message += fmt.Sprintf("<@%s> ", player)
			}
		}

		_, err := s.ChannelMessageSend(m.ChannelID, message)
		if err != nil {
			fmt.Println("Error sending cs assemble message: ", err)
		}
	}
}
