# Golang-Weather-Cli
A simple command-line tool written in Go to fetch current weather information for any city using the [OpenWeatherMap API](https://openweathermap.org/api).

---

## **Features**

- Fetch current temperature, feels like, min/max temperatures, humidity, pressure
- Weather description (clouds, rain, etc.)
- Safe API key management via `.env` file
- Simple, easy-to-use CLI

---

## **Setup**

### **1. Clone the repository**

```bash
git clone https://github.com/ginoflion/Golang-Weather-Cli.git
cd Golang-Weather-Cli
```

### **2. Create a .env file**

Create a .env file in the project root and add your OpenWeatherMap API key:
```
OPENWEATHER_API_KEY=your_api_key_here
```

### **3. Install dependencies**

This project uses godotenv to load the .env file. Install it with:
```
go get github.com/joho/godotenv
```

### **4. Run the CLI**
go run main.go <city>

## Example:
```
go run main.go London
```

## Output:
```Weather in London: Clouds
Description: broken clouds
Temperature: 13.2°C
Feels Like: 12.6°C
Min Temperature: 11.5°C
Max Temperature: 13.9°C
Humidity: 78%
Pressure: 1021 hPa
```
