package api

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lycoris11/ai-agent/internal/model"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func getYoutubeService(ctx context.Context, refreshToken string, clientID string, clientSecret string) (*youtube.Service, error) {
	// Compose config
	config := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     google.Endpoint,
		Scopes:       []string{youtube.YoutubeUploadScope},
		RedirectURL:  "http://localhost:8080", // TODO::CHANGE ME IN PRODUCTION IF API IS TURNED PUBLIC
	}
	token := &oauth2.Token{RefreshToken: refreshToken}

	ts := config.TokenSource(ctx, token)
	return youtube.NewService(ctx, option.WithTokenSource(ts))
}

func UploadVideo(c *gin.Context, google_auth *model.Google) {

	title := c.PostForm("title")
	description := c.PostForm("description")

	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"error": "No file uploaded"})
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		log.Fatalf("Error opening video file: %v", err)
		c.JSON(500, gin.H{"error": "Failed to open uploaded file"})
		return
	}
	defer file.Close()

	ctx := context.Background()
	yt, err := getYoutubeService(ctx, google_auth.GoogleRefreshToken, google_auth.ClientID, google_auth.ClientSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create YouTube service: " + err.Error()})
		return
	}

	video := &youtube.Video{
		Snippet: &youtube.VideoSnippet{
			Title:       title,
			Description: description,
			CategoryId:  "26",
			Tags:        []string{"Weather", "News", "Forecast", "AI Weather", "AI Generated", "7 day forecast"},
		},
		Status: &youtube.VideoStatus{
			PrivacyStatus: "public",
			MadeForKids:   false,
		},
	}

	call := yt.Videos.Insert([]string{"snippet", "status"}, video)
	call.Media(file, googleapi.ContentType("video/*"))

	response, err := call.Do()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "YouTube upload failed: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"videoId": response.Id, "uploadSuccess": true, "title": title})
}
