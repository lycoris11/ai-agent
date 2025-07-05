package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lycoris11/ai-agent/internal/service"
)

func SetupRouter(openAIApiKey string, weatherAIApiKey string) *gin.Engine {
	r := gin.Default()
	r.POST("/weatherScript", func(c *gin.Context) {
		service.GetAIResponse(c, openAIApiKey)
	})
	r.GET("/weather/:city", func(c *gin.Context) {
		service.GetWeatherResponse(c, weatherAIApiKey)
	})
	return r
}
