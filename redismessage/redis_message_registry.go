package redismessage

import (
	"fmt"
	"strings"

	"github.com/ShervilG/spinboard-go/discord"
	"github.com/go-redis/redis/v8"
)

func HandleRedisMessage(message *redis.Message) {
	messageString := message.String()
	fmt.Println(messageString)

	if strings.Contains(messageString, "DISCORD_REMINDER") {
		HandleReminderCallback(message)
	}
}

func HandleReminderCallback(m *redis.Message) {
	ds := discord.GetDiscordSession()

	if ds == nil {
		return
	}

	messageContent := m.String()
	parts := strings.Split(messageContent, "::")
	if len(parts) != 5 {
		return
	}

	reminderMessage := fmt.Sprintf("Here's your reminder <@%s>\n%s", parts[3], strings.TrimSuffix(parts[4], ">"))
	ds.ChannelMessageSend(parts[2], reminderMessage)
}
