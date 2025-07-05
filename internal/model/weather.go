package model

type Condition struct {
	Text string `json:"text"`
	Icon string `json:"icon"`
	Code int    `json:"code"`
}

type HourData struct {
	TimeEpoch    int       `json:"time_epoch"`
	Time         string    `json:"time"`
	TempF        float64   `json:"temp_f"`
	IsDay        int       `json:"is_day"`
	Condition    Condition `json:"condition"`
	WindMph      float64   `json:"wind_mph"`
	WindKph      float64   `json:"wind_kph"`
	WindDegree   int       `json:"wind_degree"`
	WindDir      string    `json:"wind_dir"`
	PressureMb   float64   `json:"pressure_mb"`
	PressureIn   float64   `json:"pressure_in"`
	PrecipMm     float64   `json:"precip_mm"`
	PrecipIn     float64   `json:"precip_in"`
	SnowCm       float64   `json:"snow_cm"`
	Humidity     int       `json:"humidity"`
	Cloud        int       `json:"cloud"`
	FeelslikeF   float64   `json:"feelslike_f"`
	WindchillF   float64   `json:"windchill_f"`
	HeatindexF   float64   `json:"heatindex_f"`
	DewpointF    float64   `json:"dewpoint_f"`
	WillItRain   int       `json:"will_it_rain"`
	ChanceOfRain int       `json:"chance_of_rain"`
	WillItSnow   int       `json:"will_it_snow"`
	ChanceOfSnow int       `json:"chance_of_snow"`
	VisKm        float64   `json:"vis_km"`
	VisMiles     float64   `json:"vis_miles"`
	GustMph      float64   `json:"gust_mph"`
	GustKph      float64   `json:"gust_kph"`
	UV           float64   `json:"uv"`
}

type ForecastDay struct {
	Hour []HourData `json:"hour"`
}

type Forecast struct {
	Forcastday []ForecastDay `json:"forecastDay"`
}

type WeatherResponse struct {
	Forcast Forecast `json:"forecast"`
}
