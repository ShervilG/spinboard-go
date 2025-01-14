package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/ShervilG/spinboard-go/cache"
)

type WeatherData struct {
	Status              string         `json:"status"`
	Message             string         `json:"message"`
	DeviceType          int            `json:"device_type"`
	LocalityWeatherData WeatherDetails `json:"locality_weather_data"`
}

type WeatherDetails struct {
	Temperature      float64 `json:"temperature"`
	Humidity         float64 `json:"humidity"`
	WindSpeed        float64 `json:"wind_speed"`
	WindDirection    float64 `json:"wind_direction"`
	RainIntensity    float64 `json:"rain_intensity"`
	RainAccumulation float64 `json:"rain_accumulation"`
	AqiPm10          float64 `json:"aqi_pm_10"`
	AqiPm25          float64 `json:"aqi_pm_2_point_5"`
}

func GetWeatherByLocalityId(localityId string) (*WeatherData, error) {
	var data WeatherData

	cachedValue := cache.Get(localityId)
	if cachedValue != "" {
		json.Unmarshal([]byte(cachedValue), &data)
		return &data, nil
	}

	apiToken := os.Getenv("WEATHER_DOMAIN_API_TOKEN")
	url := "https://www.weatherunion.com/gw/weather/external/v0/get_locality_weather_data?locality_id=" + localityId

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}

	req.Header.Add("X-Zomato-Api-Key", apiToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}

	// Unmarshal the JSON response into the WeatherData struct
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil, err
	}

	cache.Set(localityId, string(body), 15*time.Minute)
	return &data, nil
}
