package qweather

import (
	"encoding/json"
	"fmt"
	"time"
)

// GridWeatherNowResponse represents the response for grid weather now API
type GridWeatherNowResponse struct {
	Code       string       `json:"code"`
	UpdateTime QWeatherTime `json:"updateTime"`
	FxLink     string       `json:"fxLink"`
	Now        WeatherNow   `json:"now"`
	Refer      Reference    `json:"refer"`
}

// WeatherNow represents current weather data
type WeatherNow struct {
	ObsTime   QWeatherTime `json:"obsTime"`
	Temp      string       `json:"temp"`
	Icon      string       `json:"icon"`
	Text      string       `json:"text"`
	Wind360   string       `json:"wind360"`
	WindDir   string       `json:"windDir"`
	WindScale string       `json:"windScale"`
	WindSpeed string       `json:"windSpeed"`
	Humidity  string       `json:"humidity"`
	Precip    string       `json:"precip"`
	Pressure  string       `json:"pressure"`
	Cloud     string       `json:"cloud,omitempty"`
	Dew       string       `json:"dew,omitempty"`
}

// Reference represents data source and license information
type Reference struct {
	Sources []string `json:"sources,omitempty"`
	License []string `json:"license,omitempty"`
}

// GridWeatherNowOptions represents options for grid weather now API
type GridWeatherNowOptions struct {
	Location string // Required: longitude,latitude coordinates (e.g., "116.41,39.92")
	Lang     string // Optional: language setting
	Unit     string // Optional: unit setting (m for metric, i for imperial)
}

// WeatherNowResponse represents the response for weather now API
type WeatherNowResponse struct {
	Code       string         `json:"code"`
	UpdateTime QWeatherTime   `json:"updateTime"`
	FxLink     string         `json:"fxLink"`
	Now        WeatherNowData `json:"now"`
	Refer      Reference      `json:"refer"`
}

// WeatherNowData represents current weather data for city weather
type WeatherNowData struct {
	ObsTime   QWeatherTime `json:"obsTime"`
	Temp      string       `json:"temp"`
	FeelsLike string       `json:"feelsLike"` // 体感温度
	Icon      string       `json:"icon"`
	Text      string       `json:"text"`
	Wind360   string       `json:"wind360"`
	WindDir   string       `json:"windDir"`
	WindScale string       `json:"windScale"`
	WindSpeed string       `json:"windSpeed"`
	Humidity  string       `json:"humidity"`
	Precip    string       `json:"precip"`
	Pressure  string       `json:"pressure"`
	Vis       string       `json:"vis"` // 能见度
	Cloud     string       `json:"cloud,omitempty"`
	Dew       string       `json:"dew,omitempty"`
}

// WeatherNowOptions represents options for weather now API
type WeatherNowOptions struct {
	Location string // Required: LocationID or coordinates (e.g., "101010100" or "116.41,39.92")
	Lang     string // Optional: language setting
	Unit     string // Optional: unit setting (m for metric, i for imperial)
}

// Error represents API error response
type Error struct {
	Code    string `json:"code"`
	Message string `json:"message,omitempty"`
}

func (e Error) Error() string {
	return fmt.Sprintf("QWeather API error %s: %s", e.Code, e.Message)
}

// QWeatherTime 自定义时间类型，用于解析和风天气的时间格式
type QWeatherTime struct {
	time.Time
}

// UnmarshalJSON 实现自定义的JSON时间解析
func (t *QWeatherTime) UnmarshalJSON(data []byte) error {
	var timeStr string
	if err := json.Unmarshal(data, &timeStr); err != nil {
		return err
	}

	// 尝试多种时间格式
	formats := []string{
		"2006-01-02T15:04:05Z07:00",
		"2006-01-02T15:04:05-07:00",
		"2006-01-02T15:04+07:00",
		"2006-01-02T15:04-07:00",
		"2006-01-02T15:04:05+07:00",
		"2006-01-02T15:04:05.000Z07:00",
		time.RFC3339,
		time.RFC3339Nano,
	}

	for _, format := range formats {
		if parsedTime, err := time.Parse(format, timeStr); err == nil {
			t.Time = parsedTime
			return nil
		}
	}

	return fmt.Errorf("unable to parse time: %s", timeStr)
}

// MarshalJSON 实现自定义的JSON时间序列化
func (t QWeatherTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Time.Format(time.RFC3339))
}

// Format 便利方法，用于格式化时间
func (t QWeatherTime) Format(layout string) string {
	return t.Time.Format(layout)
}
