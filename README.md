# AI-Agent: Automated 3-Day Weather Report Generator & Uploader

This project is an **Agentic Workflow** that automates the process of generating, rendering, and uploading 3-day weather forecast videos to YouTube. It leverages weather APIs, AI-generated scripts throught GPT 4.1, dynamic image creation, and an AI video generation service to deliver 3 day weather updates with zero human intervention.

e.g. [Watch a generated weather report on YouTube!](https://www.youtube.com/watch?v=4tkDLn0-j4A)

I wrote a Go API using Gin as a middleman/proxy in between the various AI APIs I use. I was thinking long term that maybe in the future, I could give users the ability to generate their own videos. Unfortunately, generating AI videos can be costly, so I would want to add OAuth 2.0 authentication to prevent abuse of my API keys. Additionally, it gives me room to extend functionality.

I host this Agentic Worflow on EC2 and I've provided sample ec2 user data.
The Go API lives within a docker container. A cron job runs the Python script calling the Go API. Currently, the container is not exposed to public IPs. When running `docker run` only localhost is exposed. I have not yet found a reason compelling enough to make the API public. In the future when I enact rate limits, and OAuth I would consider this.

Please contribute to this using your own keys!

Weather, unlike other news segments, is completely unbiased. Its numbers, its facts, its something that lends itsself to AI! I would love help! This is open source!

## Features

- **Weather Data Fetching:** Retrieves 3-day weather forecasts for a specified city using the WeatherAPI.
- **AI Script Generation:** Uses OpenAI's GPT models to generate natural-sounding weather report scripts based on forecast data.
- **Dynamic Image Creation:** Renders background images with weather icons and data using Python and Pillow.
- **Video Generation:** Integrates with HeyGen API to create videos combining AI scripts, avatars, and generated backgrounds.
- **YouTube Upload:** Automatically uploads the final video to YouTube using the YouTube Data API.
- **Agentic Workflow:** Orchestrates the entire process via a Python cron script and a Go-based backend API.

## Project Structure

- `cmd/ai-agent/`: Main Go application entrypoint.
- `internal/api/`: API handlers and route definitions (Gin framework).
- `internal/service/`: Business logic for weather, AI, video, and image processing.
- `internal/model/`: Data models for weather, video, and API keys.
- `internal/config/`: Environment variable and API key management.
- `video_creation_cron.py`: Python script that orchestrates the daily workflow.
- `bg_image_creation.py`: Python module for generating weather background images.
- `assets/`: Fonts and generated images.
- `.env`: Environment variables (API keys, secrets, etc).

## Workflow Overview

1. **Fetch Weather Data:** The Python script requests a 3-day forecast from the Go API.
2. **Prepare Data:** Extracts and formats the next 3 days for image and script generation.
3. **Generate Background Image:** Using Python Image libraries, renders a composite image with weather icons and stats.
4. **Upload Image:** Sends the image to the Go API, which uploads it to HeyGen.
5. **Generate Script:** Sends weather data to the Go API, which uses OpenAI to generate a script.
6. **Generate Video:** Requests HeyGen to create a video using the script and background.
7. **Poll for Completion:** Waits for video generation to finish.
8. **Download Video:** Retrieves the finished video.
9. **Upload to YouTube:** Posts the video to YouTube with a generated title and description.

## Setup & Usage

### Prerequisites

- Go 1.24+
- Python 3.9.23+
- Docker (for deployment)
- API keys for OpenAI, WeatherAPI, HeyGen, and Google/YouTube

### Environment Variables

Create a `.env` file in your ec2 and reference it as an "--env-file" when running the `docker run` command:
OPENAI_API_KEY=your_openai_key
WEATHER_API_KEY=your_weatherapi_key
HEY_GEN_VIDEO_API_KEY=your_heygen_key
ENV="prod" or "dev"
REFRESH_TOKEN=your_google_refresh_token
CLIENT_ID=your_google_client_id
CLIENT_SECRET=your_google_client_secret


Note: This project is for educational and demonstration purposes. Ensure you comply with the terms of service for all third-party APIs used.