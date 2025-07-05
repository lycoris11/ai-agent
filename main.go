package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

type Condition struct {
	Text string `json:"text"`
	Icon string `json:"icon"`
	Code int    `json:"code"`
}

type HourData struct {
	TimeEpoch    int       `json:"time_epoch"`
	Time         string    `json:"time"`
	TempF        float64   `json:"temp_f"`
	IsDay        int       `json:"is_day"`
	Condition    Condition `json:"condition"`
	WindMph      float64   `json:"wind_mph"`
	WindKph      float64   `json:"wind_kph"`
	WindDegree   int       `json:"wind_degree"`
	WindDir      string    `json:"wind_dir"`
	PressureMb   float64   `json:"pressure_mb"`
	PressureIn   float64   `json:"pressure_in"`
	PrecipMm     float64   `json:"precip_mm"`
	PrecipIn     float64   `json:"precip_in"`
	SnowCm       float64   `json:"snow_cm"`
	Humidity     int       `json:"humidity"`
	Cloud        int       `json:"cloud"`
	FeelslikeF   float64   `json:"feelslike_f"`
	WindchillF   float64   `json:"windchill_f"`
	HeatindexF   float64   `json:"heatindex_f"`
	DewpointF    float64   `json:"dewpoint_f"`
	WillItRain   int       `json:"will_it_rain"`
	ChanceOfRain int       `json:"chance_of_rain"`
	WillItSnow   int       `json:"will_it_snow"`
	ChanceOfSnow int       `json:"chance_of_snow"`
	VisKm        float64   `json:"vis_km"`
	VisMiles     float64   `json:"vis_miles"`
	GustMph      float64   `json:"gust_mph"`
	GustKph      float64   `json:"gust_kph"`
	UV           float64   `json:"uv"`
}

type ForecastDay struct {
	Hour []HourData `json:"hour"`
}

type Forecast struct {
	Forcastday []ForecastDay `json:"forecastDay"`
}

type WeatherResponse struct {
	Forcast Forecast `json:"forecast"`
}

func getOpenAIAPIKey() string {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found.")
	}

	openAIApiKey := os.Getenv("OPENAI_API_KEY")
	if openAIApiKey == "" {
		log.Fatal("OPENAI_API_KEY is not set in the environment")
	}

	fmt.Println("OpenAI API Key loaded successfully.")

	return openAIApiKey
}

func getWeatherAPIKey() string {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found.")
	}

	weatherAIApiKey := os.Getenv("WEATHER_API_KEY")
	if weatherAIApiKey == "" {
		log.Fatal("WEATHER_API_KEY is not set in the environment")
	}

	fmt.Println("WeatherAPI.com API Key loaded successfully.")

	return weatherAIApiKey
}

func main() {

	openAIApiKey := getOpenAIAPIKey()
	weatherAIApiKey := getWeatherAPIKey()

	r := gin.Default()
	r.POST("/weatherScript", func(c *gin.Context) {
		getAIResponse(c, openAIApiKey)
	})
	r.GET("/weather/:city", func(c *gin.Context) {
		getWeatherResponse(c, weatherAIApiKey)
	})
	r.Run("localhost:8080")
}

func getWeatherResponse(c *gin.Context, weatherAIApiKey string) {

	city := c.Param("city")
	if city == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Missing city parameter"})
		return
	}

	url := fmt.Sprintf("https://api.weatherapi.com/v1/forecast.json?key=%s&q=%s&days=1", weatherAIApiKey, city)

	res, err := http.Get(url)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to call weather API"})
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		c.JSON(res.StatusCode, gin.H{"error": "Weather API returned non-200 status"})
		return
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read weather API response"})
		return
	}

	var weatherRes WeatherResponse
	if err := json.Unmarshal(body, &weatherRes); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse weather data"})
	}

	var hours []HourData
	if len(weatherRes.Forcast.Forcastday) > 0 {
		hours = weatherRes.Forcast.Forcastday[0].Hour
	}

	trimmed_hours := hours[7:22]

	c.IndentedJSON(http.StatusOK, trimmed_hours)
}

func getAIResponse(c *gin.Context, openAIApiKey string) {

	var request []HourData

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	jsonData, err := json.MarshalIndent(request, "", "  ")
	if err != nil {
		return
	}

	prompt := fmt.Sprintf(`Pretend you're the weatherman! Take this hourly weather data and generate a script you would read to people in the morning!:
"""
%s
"""
`, string(jsonData))

	client := openai.NewClient(
		option.WithAPIKey(openAIApiKey),
	)

	chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(prompt),
		},
		Model: openai.ChatModelGPT4o,
	})
	if err != nil {
		panic(err.Error())
	}
	println(chatCompletion.Choices[0].Message.Content)

	c.IndentedJSON(http.StatusOK, chatCompletion.Choices[0].Message.Content)
}
