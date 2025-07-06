package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lycoris11/ai-agent/internal/model"
	"github.com/lycoris11/ai-agent/internal/service"
)

func SetupRouter(keys *model.APIKeys) *gin.Engine {
	r := gin.Default()
	r.POST("/ai/weatherScript", func(c *gin.Context) {
		service.GetAIResponse(c, keys.OpenAIApiKey)
	})
	r.GET("/weather/:city", func(c *gin.Context) {
		service.GetWeatherResponse(c, keys.WeatherAPIKey)
	})
	r.POST("/ai/videoGen", func(c *gin.Context) {
		service.GetWeatherResponse(c, keys.HeyGenVideoAPIKey)
	})
	return r
}
