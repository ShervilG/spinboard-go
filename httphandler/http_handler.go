package httphandler

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/ShervilG/spinboard-go/weather"
)

func TimeHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "The time is: "+time.Now().String())
}

func PingHanlder(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Pong")
}

func WeatherHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	localityId := queryParams.Get("locality_id")
	if localityId != "" {
		weatherData, err := weather.GetWeatherByLocalityId(localityId)
		if err != nil {
			http.Error(w, "Error getting weather data", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(weatherData)
	} else {
		http.Error(w, "locality_id query param is required", http.StatusBadRequest)
	}
}
