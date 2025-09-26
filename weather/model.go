package weather

// WeatherResponse represents the JSON from OpenWeatherMap
type WeatherResponse struct {
	Main struct {
		Temp float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		MinTemp float64 `json:"temp_min"`
		MaxTemp float64 `json:"temp_max"`
		Humidity int     `json:"humidity"`
		Pressure int     `json:"pressure"`
	} `json:"main"`
	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
}

// WeatherData is the simplified struct we return
type WeatherData struct {
	Temp        float64
	Description string
	Main       string
	FeelsLike   float64
	MinTemp    float64
	MaxTemp    float64
	Humidity   int
	Pressure   int
	MainDescription string

}
