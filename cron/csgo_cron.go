package cron

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

var CSGO_CHANNEL_ID string = "711131294452678707"

func ScheduleCsgoReminderMessage(s *discordgo.Session) {
	ch := make(chan bool)

	go func() {
		for {
			currentTime := time.Now().UTC()
			istLocation, err := time.LoadLocation("Asia/Kolkata")
			if err != nil {
				continue
			}

			istTime := currentTime.In(istLocation)
			hour := istTime.Hour()
			if hour >= 20 && hour < 21 {
				if s != nil {
					s.ChannelMessageSend(CSGO_CHANNEL_ID, "Playing CS Today folks ?")
				}
			}

			time.Sleep(1 * time.Hour)
		}
	}()

	<-ch
}
