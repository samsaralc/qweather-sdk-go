package qweather

import (
	"encoding/json"
	"fmt"
	"time"
)

// GridWeatherNowResponse represents the response for grid weather now API
type GridWeatherNowResponse struct {
	Code       string       `json:"code"`       // API 状态码，"200" 表示请求成功。
	UpdateTime QWeatherTime `json:"updateTime"` // 当前 API 数据的最近更新时间。
	FxLink     string       `json:"fxLink"`     // 当前数据对应的响应式页面链接，便于在网页或应用中展示。
	Now        WeatherNow   `json:"now"`        // 当前格点实时天气数据。
	Refer      Reference    `json:"refer"`      // 数据来源和许可说明。
}

// WeatherNow represents current weather data
type WeatherNow struct {
	ObsTime   QWeatherTime `json:"obsTime"`         // 数据观测时间。格点实时天气采用 UTC 0 时区表示。
	Temp      string       `json:"temp"`            // 温度，默认单位为摄氏度。
	Icon      string       `json:"icon"`            // 天气状况图标代码，可配合和风天气图标资源使用。
	Text      string       `json:"text"`            // 天气状况文字描述，例如晴、多云、雨。
	Wind360   string       `json:"wind360"`         // 风向角度，取值范围 0-360。
	WindDir   string       `json:"windDir"`         // 风向文字描述，例如东北风。
	WindScale string       `json:"windScale"`       // 风力等级。
	WindSpeed string       `json:"windSpeed"`       // 风速，默认单位为公里/小时。
	Humidity  string       `json:"humidity"`        // 相对湿度，百分比数值。
	Precip    string       `json:"precip"`          // 过去 1 小时降水量，默认单位为毫米。
	Pressure  string       `json:"pressure"`        // 大气压强，默认单位为百帕。
	Cloud     string       `json:"cloud,omitempty"` // 云量，百分比数值，可能为空。
	Dew       string       `json:"dew,omitempty"`   // 露点温度，可能为空。
}

// Reference represents data source and license information
type Reference struct {
	Sources []string `json:"sources,omitempty"` // 原始数据来源或数据源说明，可能为空。
	License []string `json:"license,omitempty"` // 数据许可或版权声明，可能为空。
}

// GridWeatherNowOptions represents options for grid weather now API
type GridWeatherNowOptions struct {
	Location string // 查询坐标，经度和纬度使用英文逗号分隔，例如 "116.41,39.92"。
	Lang     string // 多语言设置，具体可选值参考和风天气语言参数文档。
	Unit     string // 度量衡单位，"m" 为公制，"i" 为英制。
}

// WeatherNowResponse represents the response for weather now API
type WeatherNowResponse struct {
	Code       string         `json:"code"`       // API 状态码，"200" 表示请求成功。
	UpdateTime QWeatherTime   `json:"updateTime"` // 当前 API 数据的最近更新时间。
	FxLink     string         `json:"fxLink"`     // 当前数据对应的响应式页面链接，便于在网页或应用中展示。
	Now        WeatherNowData `json:"now"`        // 当前城市实时天气数据。
	Refer      Reference      `json:"refer"`      // 数据来源和许可说明。
}

// WeatherNowData represents current weather data for city weather
type WeatherNowData struct {
	ObsTime   QWeatherTime `json:"obsTime"`         // 数据观测时间。
	Temp      string       `json:"temp"`            // 温度，默认单位为摄氏度。
	FeelsLike string       `json:"feelsLike"`       // 体感温度，默认单位为摄氏度。
	Icon      string       `json:"icon"`            // 天气状况图标代码，可配合和风天气图标资源使用。
	Text      string       `json:"text"`            // 天气状况文字描述，例如晴、多云、雨。
	Wind360   string       `json:"wind360"`         // 风向角度，取值范围 0-360。
	WindDir   string       `json:"windDir"`         // 风向文字描述，例如东南风。
	WindScale string       `json:"windScale"`       // 风力等级。
	WindSpeed string       `json:"windSpeed"`       // 风速，默认单位为公里/小时。
	Humidity  string       `json:"humidity"`        // 相对湿度，百分比数值。
	Precip    string       `json:"precip"`          // 过去 1 小时降水量，默认单位为毫米。
	Pressure  string       `json:"pressure"`        // 大气压强，默认单位为百帕。
	Vis       string       `json:"vis"`             // 能见度，默认单位为公里。
	Cloud     string       `json:"cloud,omitempty"` // 云量，百分比数值，可能为空。
	Dew       string       `json:"dew,omitempty"`   // 露点温度，可能为空。
}

// WeatherNowOptions represents options for weather now API
type WeatherNowOptions struct {
	Location string // 查询地区，可传 LocationID 或 "经度,纬度" 坐标，例如 "101010100" 或 "116.41,39.92"。
	Lang     string // 多语言设置，具体可选值参考和风天气语言参数文档。
	Unit     string // 度量衡单位，"m" 为公制，"i" 为英制。
}

// Error represents API error response
type Error struct {
	Code    string `json:"code"`              // 和风天气 API 返回的状态码。
	Message string `json:"message,omitempty"` // 对状态码的补充错误描述。
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
