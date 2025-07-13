package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lycoris11/ai-agent/internal/model"
	"github.com/lycoris11/ai-agent/internal/service"
)

func SetupRouter(keys *model.APIKeys, google_auth *model.Google) *gin.Engine {
	r := gin.Default()
	r.GET("/weather/hourly/:city", func(c *gin.Context) {
		service.GetHourlyWeatherResponse(c, keys.WeatherAPIKey)
	})
	r.GET("/weather/7day/:city", func(c *gin.Context) {
		service.Get7DayWeatherResponse(c, keys.WeatherAPIKey)
	})

	r.POST("/ai/hourly/weatherScript", func(c *gin.Context) {
		service.GetHourlyAIResponse(c, keys.OpenAIApiKey)
	})
	r.POST("/ai/7day/weatherScript", func(c *gin.Context) {
		service.Get7DayAIResponse(c, keys.OpenAIApiKey)
	})

	r.POST("/video/backgroundImageUpload", func(c *gin.Context) {
		service.UploadImage(c, keys.HeyGenVideoAPIKey)
	})
	r.POST("/video/generateVideo", func(c *gin.Context) {
		service.GenerateAIVideo(c, keys.HeyGenVideoAPIKey)
	})
	r.GET("/video/getStatus", func(c *gin.Context) {
		service.GetStatus(c, keys.HeyGenVideoAPIKey)
	})

	r.POST("/uploadVideo", func(c *gin.Context) {
		UploadVideo(c, google_auth)
	})
	return r
}
