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
		log.Println("No .env file or system environment variable found.")
	}

	weatherAIApiKey := os.Getenv("WEATHER_API_KEY")
	if weatherAIApiKey == "" {
		log.Fatal("WEATHER_API_KEY is not set in the environment")
	}

	fmt.Println("WeatherAPI.com API Key loaded successfully.")

	return weatherAIApiKey
}

func GetHeyGenVideoAPIKey() string {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file or system environment variable found.")
	}

	heyGenVideoAPIKey := os.Getenv("HEY_GEN_VIDEO_API_KEY")
	if heyGenVideoAPIKey == "" {
		log.Fatal("HEY_GEN_VIDEO_API_KEY is not set in the environment")
	}

	fmt.Println("HEY_GEN_VIDEO_API_KEY API Key loaded successfully.")

	return heyGenVideoAPIKey
}

func GetEnv() string {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file or system environment variable found.")
	}

	env := os.Getenv("ENV")
	if env == "" {
		log.Fatal("ENV is not set in the environment")
	}

	fmt.Println("Environment loaded successfully.")

	return env
}
