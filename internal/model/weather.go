package model

type Condition struct {
	Text string `json:"text"`
	Icon string `json:"icon"`
	Code int    `json:"code"`
}

type HourData struct {
	TimeEpoch    int       `json:"time_epoch"`
	Time         string    `json:"time"`
	TempF        float32   `json:"temp_f"`
	IsDay        int       `json:"is_day"`
	Condition    Condition `json:"condition"`
	WindMph      float32   `json:"wind_mph"`
	WindKph      float32   `json:"wind_kph"`
	WindDegree   int       `json:"wind_degree"`
	WindDir      string    `json:"wind_dir"`
	PressureMb   float32   `json:"pressure_mb"`
	PressureIn   float32   `json:"pressure_in"`
	PrecipMm     float32   `json:"precip_mm"`
	PrecipIn     float32   `json:"precip_in"`
	SnowCm       float32   `json:"snow_cm"`
	Humidity     int       `json:"humidity"`
	Cloud        int       `json:"cloud"`
	FeelslikeF   float32   `json:"feelslike_f"`
	WindchillF   float32   `json:"windchill_f"`
	HeatindexF   float32   `json:"heatindex_f"`
	DewpointF    float32   `json:"dewpoint_f"`
	WillItRain   int       `json:"will_it_rain"`
	ChanceOfRain int       `json:"chance_of_rain"`
	WillItSnow   int       `json:"will_it_snow"`
	ChanceOfSnow int       `json:"chance_of_snow"`
	VisKm        float32   `json:"vis_km"`
	VisMiles     float32   `json:"vis_miles"`
	GustMph      float32   `json:"gust_mph"`
	GustKph      float32   `json:"gust_kph"`
	UV           float32   `json:"uv"`
}

type Hourly_ForecastDay struct {
	Hour []HourData `json:"hour"`
}

type Hourly_Forecast struct {
	Forcastday []Hourly_ForecastDay `json:"forecastday"`
}

type Hourly_WeatherResponse struct {
	Forcast Hourly_Forecast `json:"forecast"`
}

type SevenDay_DayData struct {
	Maxtemp_f            float32   `json:"maxtemp_f"`
	Mintemp_f            float32   `json:"mintemp_f"`
	Avgtemp_f            float32   `json:"avgtemp_f"`
	Maxwind_mph          float32   `json:"maxwind_mph"`
	Totalprecip_in       float32   `json:"totalprecip_in"`
	Totalsnow_cm         float32   `json:"totalsnow_cm"`
	Avgvis_miles         float32   `json:"avgvis_miles"`
	Avghumidity          float32   `json:"avghumidity"`
	Daily_will_it_rain   int       `json:"daily_will_it_rain"`
	Daily_chance_of_rain int       `json:"daily_chance_of_rain"`
	Daily_will_it_snow   int       `json:"daily_will_it_snow"`
	Daily_chance_of_snow int       `json:"daily_chance_of_snow"`
	Condition            Condition `json:"condition"`
	UV                   float32   `json:"uv"`
}

type SevenDay_ForecastDay struct {
	Date string           `json:"date"`
	Day  SevenDay_DayData `json:"day"`
}

type SevenDay_Forecast struct {
	Forcastday []SevenDay_ForecastDay `json:"forecastday"`
}

type SevenDay_WeatherResponse struct {
	Forcast SevenDay_Forecast `json:"forecast"`
}
