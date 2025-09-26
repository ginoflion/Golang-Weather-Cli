package main

import (
	"fmt"
	"os"
	"github.com/ginoflion/weather-cli/weather"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()	
	
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	if len(os.Args) < 2 {
		fmt.Println("Usage: weather <city>")
		return
	}

	city := os.Args[1]

	data, err := weather.GetCurrentWeather(city)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Weather in %s: %s\n", city, data.MainDescription)
	fmt.Printf("Description: %s\n", data.Description)
	fmt.Printf("Temperature: %.1f°C\n", data.Temp)
	fmt.Printf("Feels Like: %.1f°C\n", data.FeelsLike)
	fmt.Printf("Min Temperature: %.1f°C\n", data.MinTemp)
	fmt.Printf("Max Temperature: %.1f°C\n", data.MaxTemp)
	fmt.Printf("Humidity: %d%%\n", data.Humidity)
	fmt.Printf("Pressure: %d hPa\n", data.Pressure)
}
