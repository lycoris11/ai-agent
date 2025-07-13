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

func GetGoogleRefreshToken() string {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file or system environment variable found.")
	}

	REFRESH_TOKEN := os.Getenv("REFRESH_TOKEN")
	if REFRESH_TOKEN == "" {
		log.Fatal("REFRESH_TOKEN is not set in the environment")
	}

	fmt.Println("REFRESH_TOKEN loaded successfully.")

	return REFRESH_TOKEN
}

func GetGoogleClientID() string {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file or system environment variable found.")
	}

	CLIENT_ID := os.Getenv("CLIENT_ID")
	if CLIENT_ID == "" {
		log.Fatal("CLIENT_ID is not set in the environment")
	}

	fmt.Println("CLIENT_ID loaded successfully.")

	return CLIENT_ID
}

func GetGoogleClientSecret() string {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file or system environment variable found.")
	}

	CLIENT_SECRET := os.Getenv("CLIENT_SECRET")
	if CLIENT_SECRET == "" {
		log.Fatal("CLIENT_SECRET is not set in the environment")
	}

	fmt.Println("CLIENT_SECRET loaded successfully.")

	return CLIENT_SECRET
}
