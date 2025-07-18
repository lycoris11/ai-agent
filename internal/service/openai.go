package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lycoris11/ai-agent/internal/model"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

func GetHourlyAIResponse(c *gin.Context, openAIApiKey string) {

	var request []model.HourData

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	jsonData, err := json.Marshal(request)
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

func Get7DayAIResponse(c *gin.Context, openAIApiKey string) {
	var request []model.SevenDay_ForecastDay

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	jsonData, err := json.MarshalIndent(request, "", "  ")
	if err != nil {
		return
	}

	prompt := fmt.Sprintf(`Pretend you're the weatherman! Your name is Gale Walker. Take this 3 day weather forecast and generate a script you would read to people! This script will be read by AI so only include english words. Don't include any Intro our outro, don't have any \n in the response.:
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
