package model

type Google struct {
	GoogleRefreshToken string
	ClientID           string
	ClientSecret       string
}

type APIKeys struct {
	OpenAIApiKey      string
	WeatherAPIKey     string
	HeyGenVideoAPIKey string
	Google            Google
}
