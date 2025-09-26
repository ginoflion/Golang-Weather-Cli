package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const baseURL = "https://api.openweathermap.org/data/2.5/weather"

func GetCurrentWeather(city string) (*WeatherData, error) {
	apiKey := os.Getenv("OPENWEATHER_API_KEY")
	if apiKey == "" {
    return nil, fmt.Errorf("API key not set. Please set OPENWEATHER_API_KEY environment variable")
	}
	url := fmt.Sprintf("%s?q=%s&appid=%s&units=metric", baseURL, city, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	var wResp WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&wResp); err != nil {
		return nil, err
	}

	data := &WeatherData{
		MainDescription: wResp.Weather[0].Main,
		Temp:        wResp.Main.Temp,
		FeelsLike:   wResp.Main.FeelsLike,
		MinTemp:    wResp.Main.MinTemp,
		MaxTemp:    wResp.Main.MaxTemp,
		Humidity:   wResp.Main.Humidity,
		Pressure:   wResp.Main.Pressure,
		Description: wResp.Weather[0].Description,

	}

	return data, nil
}
