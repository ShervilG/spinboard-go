package discordhandler

import (
	"time"

	"github.com/ShervilG/spinboard-go/cache"
	"github.com/bwmarrin/discordgo"
)

var LETS_PLAY_CS_MESAGE_SENT_KEY string = "LETS_PLAY_CS_MESAGE_SENT"
var CS_VOICE_CHANNEL_ID string = "735022555240464415"

func HandleCsVoiceChannelJoin(s *discordgo.Session, vsu *discordgo.VoiceStateUpdate) {
	if cache.Get(LETS_PLAY_CS_MESAGE_SENT_KEY) != "" {
		return
	}

	if s.State.User.ID == vsu.UserID {
		return
	}

	if vsu.BeforeUpdate == nil && vsu.ChannelID == CS_VOICE_CHANNEL_ID {
		ch, err := s.Channel(vsu.ChannelID)
		if err != nil {
			return
		}

		guild, err := s.State.Guild(ch.GuildID)
		if err != nil {
			return
		}

		activeMemberCount := 0
		for _, vs := range guild.VoiceStates {
			if vs.ChannelID == CS_VOICE_CHANNEL_ID {
				activeMemberCount++
			}
		}

		if activeMemberCount >= 5 {
			s.ChannelMessageSend(ch.ID, "Lets play some CS ! 🔥")
		}

		cache.Set(LETS_PLAY_CS_MESAGE_SENT_KEY, "true", time.Minute*5)
	}
}
