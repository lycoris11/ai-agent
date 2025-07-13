package main

import (
	"github.com/lycoris11/ai-agent/internal/api"
	"github.com/lycoris11/ai-agent/internal/config"
	"github.com/lycoris11/ai-agent/internal/model"
)

func main() {
	keys := model.APIKeys{
		OpenAIApiKey:      config.GetOpenAIAPIKey(),
		WeatherAPIKey:     config.GetWeatherAPIKey(),
		HeyGenVideoAPIKey: config.GetHeyGenVideoAPIKey(),
	}
	ENV := config.GetEnv()
	google_auth := model.Google{
		GoogleRefreshToken: config.GetGoogleRefreshToken(),
		ClientID:           config.GetGoogleClientID(),
		ClientSecret:       config.GetGoogleClientSecret(),
	}

	r := api.SetupRouter(&keys, &google_auth)

	switch ENV {
	case "prod":
		r.Run(":8080")
	case "dev":
		r.Run("localhost:8080")
	}
}
