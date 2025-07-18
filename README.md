
# ⛅️ AI-Agent: Automated 3-Day Weather Report Generator & Uploader 📺

🚀 **Zero-human-intervention 3-day weather report videos, automatically generated and posted to YouTube!**  
From API to scriptwriting, image & video generation, to upload — all handled by an agentic workflow.

[▶️ Watch an auto-generated weather report!](https://www.youtube.com/watch?v=4tkDLn0-j4A)

---

## 🧩 Project Overview

This project features:
- **Agentic Workflow:** Python + Go orchestrate the show
- **AI-powered scripts:** GPT-4.1 generates natural reports
- **Dynamic images:** Visuals generated on the fly
- **AI Video:** HeyGen avatars narrate your weather update
- **Source:** WeatherAPI delivers unbiased data
- **Delivery:** Automatic upload to YouTube

Built for extensibility and security:  
- 🟦 **Go API** (with Gin) as a smart proxy and control point  
- 🟩 **Dockerized** backend, private by default  
- 🟧 **EC2-hosted workflow** with ready-to-use launch scripts  
- 🟪 **OAuth 2.0 & rate limiting planned** for future public access & security

> **🌱 Open Source!** Feel free to contribute (with your own API keys). Let’s build the future of AI-generated media together!

---

## ✨ Features at a Glance

- 🌤️ **Weather Data Fetching** — 3-day forecasts with WeatherAPI  
- 🤖 **AI Script Generation** — Natural language via OpenAI GPT  
- 🎨 **Dynamic Image Creation** — Python & Pillow composites  
- 🎭 **AI Video Production** — Avatars & backgrounds via HeyGen  
- ⬆️ **YouTube Auto-upload** — Fully automated with titles & description  
- 🔄 **End-to-End Automation** — Cronjob + Docker + Go API + Python

---

## 📂 Project Structure

| Path / File                | Purpose                                         |
|----------------------------|------------------------------------------------|
| `cmd/ai-agent/`            | Go app entrypoint                              |
| `internal/api/`            | Gin HTTP handlers                              |
| `internal/service/`        | Core business logic                            |
| `internal/model/`          | Data models (weather/video/API keys)           |
| `internal/config/`         | Env & key management                           |
| `video_creation_cron.py`   | Daily workflow orchestrator (Python)           |
| `bg_image_creation.py`     | Weather image generator (Python)               |
| `assets/`                  | Fonts & image storage                          |
| `.env`                     | Store API keys and config                      |

---

## ⚙️ How it Works

1. **Fetch Weather** 🛰️  
   Python script calls Go API for 3-day forecast
2. **Prepare Data** 📅  
   Extract days, parse for reporting
3. **Generate Image** 🖼️  
   Create weather backgrounds with Python & Pillow
4. **Image Upload** ☁️  
   Image sent to Go API → HeyGen
5. **Script Generation** 📑  
   AI-written script by OpenAI (Go API triggers)
6. **Video Creation** 🎥  
   HeyGen puts it all together: voice, avatar, background
7. **Wait & Retrieve** ⏰  
   Script polls status, downloads when ready
8. **YouTube Upload** 🚀  
   Auto-post with catchy title + description

---

## 🔧 Installation & Setup

### Prerequisites

- 🟠 **Go** v1.24+
- 🐍 **Python** 3.9.23+
- 🐋 **Docker** (for seamless deployment)
- 🔑 API keys: OpenAI, WeatherAPI, HeyGen, & Google/YouTube

### .env Example

Set these in your `.env` (and pass as `--env-file` when launching Docker):

```dotenv
OPENAI_API_KEY=your_openai_key
WEATHER_API_KEY=your_weatherapi_key
HEY_GEN_VIDEO_API_KEY=your_heygen_key
ENV=prod   # or "dev"
REFRESH_TOKEN=your_google_refresh_token
CLIENT_ID=your_google_client_id
CLIENT_SECRET=your_google_client_secret
```

---

### 🚦 Notes & Contribution

- 📚 *Educational / demonstration purposes only.*
- 🖊️ *Comply with third-party API terms of service!*
- 🤝 *PRs, Issues, and Suggestions are very welcome!*
- 🌍 *Weather is unbiased: just numbers and facts — perfect for AI automation!*

---

✨**Let’s change how news is made!**✨

---