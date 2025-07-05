package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lycoris11/ai-agent/internal/model"
)

func GetWeatherResponse(c *gin.Context, weatherAIApiKey string) {

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

	var weatherRes model.WeatherResponse
	if err := json.Unmarshal(body, &weatherRes); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse weather data"})
	}

	var hours []model.HourData
	if len(weatherRes.Forcast.Forcastday) > 0 {
		hours = weatherRes.Forcast.Forcastday[0].Hour
	}

	trimmed_hours := hours[7:22]

	c.IndentedJSON(http.StatusOK, trimmed_hours)
}
