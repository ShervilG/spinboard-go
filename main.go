package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ShervilG/spinboard-go/cache"
	"github.com/ShervilG/spinboard-go/cron"
	"github.com/ShervilG/spinboard-go/discordhandler"
	"github.com/ShervilG/spinboard-go/httphandler"
	"github.com/ShervilG/spinboard-go/rediscache"
	"github.com/bwmarrin/discordgo"
)

var ds *discordgo.Session

func main() {
	// Discord
	discordBotToken := os.Getenv("BUNTY_BOT_TOKEN")
	ds, _ = discordgo.New("Bot " + discordBotToken)

	ds.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMessages | discordgo.IntentsDirectMessages | discordgo.IntentsGuildPresences | discordgo.IntentsGuildVoiceStates
	ds.AddHandler(discordhandler.HandleHello)
	ds.AddHandler(discordhandler.HandleCsAssemble)
	ds.AddHandler(discordhandler.HandleWeather)
	ds.AddHandler(discordhandler.HandleCsVoiceChannelJoin)
	ds.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		s.StateEnabled = true
	})
	err := ds.Open()
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
	}

	defer ds.Close()

	// Crons
	scheduleCrons()

	// Setup Cache
	cache.SetupCache()
	rediscache.SetupRedisClient()

	// HTTP Server
	http.HandleFunc("/", httphandler.PingHanlder)
	http.HandleFunc("/ping", httphandler.PingHanlder)
	http.HandleFunc("/time", httphandler.TimeHandler)
	http.HandleFunc("/weather", httphandler.WeatherHandler)
	http.HandleFunc("/cache/set", httphandler.CacheSetHandler)
	http.HandleFunc("/cache/get", httphandler.CacheGetHandler)

	http.ListenAndServe(":8000", nil)
}

func scheduleCrons() {
	cron.ScheduleCsgoReminderMessage(ds)
}
