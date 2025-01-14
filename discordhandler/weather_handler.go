package discordhandler

import (
	"fmt"
	"strings"

	"github.com/ShervilG/spinboard-go/weather"
	"github.com/bwmarrin/discordgo"
)

func HandleWeather(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(m.Content, "!weather") {
		parts := strings.Split(m.Content, " ")
		if len(parts) == 2 {
			localityId := parts[1]
			weatherData, err := weather.GetWeatherByLocalityId(localityId)
			if err != nil {
				s.ChannelMessageSend(m.ChannelID, "Error getting weather data")
				return
			}

			messageString := fmt.Sprintf("Temperature: %f\nHumidity: %f\nRain Intensity: %f", weatherData.LocalityWeatherData.Temperature, weatherData.LocalityWeatherData.Humidity, weatherData.LocalityWeatherData.RainIntensity)
			s.ChannelMessageSend(m.ChannelID, messageString)
		}
	}
}
