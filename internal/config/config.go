package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetOpenAIAPIKey() string {
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

func GetWeatherAPIKey() string {
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
