package discordhandler

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/ShervilG/spinboard-go/rediscache"
	"github.com/bwmarrin/discordgo"
)

const REMINDER_KEY_PREFIX = "DISCORD_REMINDER"

func HandleReminderSet(s *discordgo.Session, m *discordgo.MessageCreate) {
	if !strings.HasPrefix(m.Content, "!rm") || s.State.User.ID == m.Author.ID {
		return
	}

	parts := strings.Split(m.Content, " ")
	if len(parts) < 3 {
		s.ChannelMessageSend(m.ChannelID, "Error in command !")
		return
	}

	message := ""
	for _, part := range parts[1 : len(parts)-1] {
		message += (" " + part)
	}
	expiry, err := strconv.ParseInt(parts[len(parts)-1], 10, 32)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Error in command !")
		return
	}

	r := rediscache.GetRedisClient()
	key := fmt.Sprintf("%s::%v::%v::%v::%v", REMINDER_KEY_PREFIX, m.GuildID, m.ChannelID, m.Author.ID, message)
	r.Set(context.Background(), key, "", time.Second*time.Duration(expiry))

	s.ChannelMessageSend(m.ChannelID, "Reminder saved")
}
