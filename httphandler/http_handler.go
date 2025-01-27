package httphandler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/ShervilG/spinboard-go/rediscache"
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

func CacheSetHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	key := queryParams.Get("key")
	val := queryParams.Get("val")

	if key != "" && val != "" {
		cacheClient := rediscache.GetRedisClient()
		cmd := cacheClient.Set(r.Context(), key, val, 5*time.Minute)
		res, err := cmd.Result()
		if err != nil {
			json.NewEncoder(w).Encode(err)
		} else {
			json.NewEncoder(w).Encode(res)
		}
	} else {
		http.Error(w, "Params not provided !", http.StatusBadRequest)
	}
}

func CacheGetHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	key := queryParams.Get("key")

	if key != "" {
		cacheClient := rediscache.GetRedisClient()
		cmd := cacheClient.Get(r.Context(), key)
		res, err := cmd.Result()
		if err != nil {
			fmt.Printf("Error occured while getting key %v\n", key)
			json.NewEncoder(w).Encode(err)
		} else {
			json.NewEncoder(w).Encode(res)
		}
	} else {
		http.Error(w, "Params not provided !", http.StatusBadRequest)
	}
}
