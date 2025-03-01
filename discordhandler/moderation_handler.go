package discordhandler

import (
	"fmt"

	"github.com/ShervilG/spinboard-go/llm"
	"github.com/bwmarrin/discordgo"
)

func HandleAndModerateAllMessages(s *discordgo.Session, m *discordgo.MessageCreate) {
	if s.State.User.ID == m.Author.ID {
		return
	}

	content := m.Content
	prompt := fmt.Sprintf(`
		Assume you are a discord moderator and you have to moderate the following message:
		"%v"
		Give a polite response to the user like: Hey that is not a good thing to say :/
		Just give a single line response, no pre text or anything. Only respond if the message is offensive/contained any cuss words otherwise reply: NULL.
	`, content)

	data := llm.GetCompletionsResponse(prompt)
	if data != "NULL" && len(data) > 0 {
		s.ChannelMessageSend(m.ChannelID, data)
	}
}
