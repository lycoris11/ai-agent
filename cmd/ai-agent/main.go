package main

import (
	"github.com/lycoris11/ai-agent/internal/api"
	"github.com/lycoris11/ai-agent/internal/config"
)

func main() {
	openAIApiKey := config.GetOpenAIAPIKey()
	weatherAPIKey := config.GetWeatherAPIKey()

	r := api.SetupRouter(openAIApiKey, weatherAPIKey)
	r.Run("localhost:8080")
}
