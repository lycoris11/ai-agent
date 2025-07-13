package model

type Background struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

type ElevenLabsSettings struct {
	Model            string  `json:"model"`
	Similarity_boost float32 `json:"similarity_boost"`
	Stability        float32 `json:"stability"`
	Style            float32 `json:"style"`
}

type Voice struct {
	Type               string             `json:"type"`
	VoiceID            string             `json:"voice_id"`
	InputText          string             `json:"input_text"`
	Speed              float32            `json:"speed"`
	ElevenLabsSettings ElevenLabsSettings `json:"elevenlabs_settings"`
}

type Offset struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

type Character struct {
	Type         string  `json:"type"`
	AvatarID     string  `json:"avatar_id"`
	Offset       Offset  `json:"offset"`
	Scale        float32 `json:"scale"`
	TalkingStyle string  `json:"talking_style"`
	Expression   string  `json:"expression"`
}

type VideoInputs struct {
	Character  Character  `json:"character"`
	Voice      Voice      `json:"voice"`
	Background Background `json:"background"`
}

type Dimension struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type VideoData struct {
	Caption     bool          `json:"caption"`
	Dimention   Dimension     `json:"dimension"`
	VideoInputs []VideoInputs `json:"video_inputs"`
}
