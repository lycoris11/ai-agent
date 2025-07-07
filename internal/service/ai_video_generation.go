package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lycoris11/ai-agent/internal/model"
)

func GenerateAIVideo(c *gin.Context, videoAPIKey string) {

	var request []model.VideoData

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	jsonPayload, err := json.Marshal(request)
	if err != nil {
		return
	}

	payload := bytes.NewBuffer(jsonPayload)

	url := "https://api.heygen.com/v2/video/generate"

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("x-api-key", videoAPIKey)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
}

func UploadImage(c *gin.Context, videoAPIKey string) {

	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get file"})
		return
	}
	defer file.Close()

	fileBytes := bytes.NewBuffer(nil)
	if _, err := io.Copy(fileBytes, file); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file"})
		return
	}

	url := "https://upload.heygen.com/v1/asset"
	req, err := http.NewRequest("POST", url, fileBytes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}

	req.Header.Add("content-type", "image/jpeg")
	req.Header.Add("X-API-KEY", "<your-api-key>")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send request"})
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response"})
		return
	}

	c.Data(res.StatusCode, res.Header.Get("Content-Type"), body)
}
