# QWeather SDK for Go

这是一个用于访问和风天气API的Go SDK，目前支持格点实时天气数据获取。

## 功能特性

- ✅ 格点实时天气数据获取
- ✅ 支持多语言设置
- ✅ 支持公制/英制单位转换
- ✅ 自动JWT身份认证
- ✅ 完整的错误处理
- ✅ 输入参数验证
- ✅ 支持GZIP压缩

## 安装

```bash
go get github.com/samsaralc/qweather-sdk-go
```

## 快速开始

### 1. 获取认证凭据

首先需要在[和风天气开发者平台](https://dev.qweather.com/)注册账号并获取认证凭据。和风天气支持两种认证方式：

- **JWT (推荐)**: 更安全的身份认证方式，需要生成JWT Token
- **API Key**: 传统的认证方式，操作简单但安全性较低

### 2. 基础使用

#### 使用JWT认证（推荐）

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/samsaralc/qweather-sdk-go"
)

func main() {
    // 方法1: 使用便利构造函数
    client := qweather.NewClientWithJWT("your_jwt_token_here")
    
    // 方法2: 使用完整配置
    config := qweather.Config{
        Token:    "your_jwt_token_here",
        AuthType: qweather.AuthTypeJWT,
        BaseURL:  "https://devapi.qweather.com", // 可选
        Timeout:  30 * time.Second, // 可选
    }
    client = qweather.NewClient(config)

    // 获取北京的格点天气数据
    location := "116.41,39.92"
    response, err := client.GetGridWeatherNow(location)
    if err != nil {
        log.Fatalf("获取天气数据失败: %v", err)
    }

    fmt.Printf("当前温度: %s°C\n", response.Now.Temp)
    fmt.Printf("天气状况: %s\n", response.Now.Text)
}
```

#### 使用API Key认证

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/samsaralc/qweather-sdk-go"
)

func main() {
    // 方法1: 使用便利构造函数
    client := qweather.NewClientWithAPIKey("your_api_key_here")
    
    // 方法2: 使用完整配置
    config := qweather.Config{
        Token:    "your_api_key_here",
        AuthType: qweather.AuthTypeAPIKey,
    }
    client = qweather.NewClient(config)

    // 获取天气数据
    location := "116.41,39.92"
    response, err := client.GetGridWeatherNow(location)
    if err != nil {
        log.Fatalf("获取天气数据失败: %v", err)
    }

    fmt.Printf("当前温度: %s°C\n", response.Now.Temp)
}
```

## API 文档

### Client 配置

```go
// 认证类型
type AuthType int

const (
    AuthTypeJWT AuthType = iota // JWT Bearer Token (推荐)
    AuthTypeAPIKey              // API KEY (传统方式)
)

type Config struct {
    BaseURL  string        // API基础URL，默认: "https://devapi.qweather.com"
    Token    string        // JWT Token 或 API Key (必需)
    AuthType AuthType      // 认证方式：JWT 或 API Key，默认为JWT
    Timeout  time.Duration // HTTP请求超时时间，默认: 30秒
}
```

#### 便利构造函数

```go
// 使用JWT认证创建客户端
func NewClientWithJWT(jwtToken string) *Client

// 使用API Key认证创建客户端  
func NewClientWithAPIKey(apiKey string) *Client

// 使用完整配置创建客户端
func NewClient(config Config) *Client
```

### 格点实时天气

#### GetGridWeatherNow

获取指定坐标的格点实时天气数据。

```go
func (c *Client) GetGridWeatherNow(location string, options ...GridWeatherNowOptions) (*GridWeatherNowResponse, error)
```

**参数:**
- `location`: 经纬度坐标字符串，格式为 "经度,纬度"，例如 "116.41,39.92"
- `options`: 可选参数，包括语言和单位设置

**返回:**
- `GridWeatherNowResponse`: 天气数据响应
- `error`: 错误信息

#### GetGridWeatherNowWithCoordinates

使用分离的经纬度数值获取格点实时天气数据的便利方法。

```go
func (c *Client) GetGridWeatherNowWithCoordinates(longitude, latitude float64, options ...GridWeatherNowOptions) (*GridWeatherNowResponse, error)
```

**参数:**
- `longitude`: 经度
- `latitude`: 纬度
- `options`: 可选参数

### 可选参数

```go
type GridWeatherNowOptions struct {
    Location string // 经纬度坐标
    Lang     string // 语言设置，例如: "zh", "en"
    Unit     string // 单位设置: "m"(公制) 或 "i"(英制)
}
```

### 响应数据结构

```go
type GridWeatherNowResponse struct {
    Code       string    `json:"code"`       // 状态码
    UpdateTime time.Time `json:"updateTime"` // API更新时间
    FxLink     string    `json:"fxLink"`     // 响应式页面链接
    Now        WeatherNow `json:"now"`       // 当前天气数据
    Refer      Reference  `json:"refer"`     // 数据源信息
}

type WeatherNow struct {
    ObsTime    time.Time `json:"obsTime"`    // 观测时间
    Temp       string    `json:"temp"`       // 温度
    Icon       string    `json:"icon"`       // 天气图标代码
    Text       string    `json:"text"`       // 天气状况描述
    Wind360    string    `json:"wind360"`    // 风向角度
    WindDir    string    `json:"windDir"`    // 风向
    WindScale  string    `json:"windScale"`  // 风力等级
    WindSpeed  string    `json:"windSpeed"`  // 风速
    Humidity   string    `json:"humidity"`   // 相对湿度
    Precip     string    `json:"precip"`     // 降水量
    Pressure   string    `json:"pressure"`   // 大气压强
    Cloud      string    `json:"cloud"`      // 云量(可能为空)
    Dew        string    `json:"dew"`        // 露点温度(可能为空)
}
```

## 使用示例

### 基础调用

```go
// 使用位置字符串
location := "116.41,39.92"
response, err := client.GetGridWeatherNow(location)
```

### 使用经纬度数值

```go
// 使用分离的经纬度
longitude := 116.41
latitude := 39.92
response, err := client.GetGridWeatherNowWithCoordinates(longitude, latitude)
```

### 使用可选参数

```go
// 设置语言和单位
options := qweather.GridWeatherNowOptions{
    Lang: "en", // 英文
    Unit: "i",  // 英制单位
}
response, err := client.GetGridWeatherNow(location, options)
```

### 错误处理

```go
response, err := client.GetGridWeatherNow(location)
if err != nil {
    // 检查是否是API错误
    if apiErr, ok := err.(qweather.Error); ok {
        fmt.Printf("API错误 %s: %s\n", apiErr.Code, apiErr.Message)
    } else {
        fmt.Printf("其他错误: %v\n", err)
    }
    return
}

// 检查API状态码
if response.Code != "200" {
    fmt.Printf("API返回非正常状态码: %s\n", response.Code)
    return
}
```

## 身份认证

### JWT认证（推荐）

JWT（JSON Web Token）是更安全的身份认证方式，使用Ed25519算法进行签名。详细的JWT生成方法请参考[和风天气JWT身份认证文档](https://dev.qweather.com/docs/resource/auth/#json-web-token)。

```go
// 使用JWT Token
client := qweather.NewClientWithJWT("your.jwt.token")
```

### API Key认证

API Key是传统的认证方式，操作简单。可以在[控制台-项目管理](https://console.qweather.com/project)中获取。

```go
// 使用API Key
client := qweather.NewClientWithAPIKey("your_api_key")
```

**注意**: 为了提高安全性，建议优先使用JWT认证方式。从2027年1月1日起，API KEY的每日请求数量将受到限制。

## 注意事项

1. **认证凭据**: 使用前请确保已设置有效的JWT Token或API Key
2. **认证方式**: 推荐使用JWT认证以获得更高的安全性
3. **坐标格式**: 经纬度坐标支持最多2位小数，格式为 "经度,纬度"
4. **时区**: 格点天气数据采用UTC 0时区表示时间
5. **数据来源**: 格点天气基于数值预报模型，不适宜与观测站数据对比
6. **分辨率**: 格点天气数据分辨率为3-5公里

## 状态码说明

常见的API状态码:
- `200`: 请求成功
- `204`: 请求成功，但你查询的地区暂时没有你需要的数据
- `400`: 请求错误，可能包含错误的请求参数或缺少必需的参数
- `401`: 认证失败，可能使用了错误的KEY、数字签名错误、KEY的类型错误等
- `403`: 无访问权限，可能是绑定的PackageName、BundleID、域名IP地址不一致，或者是需要额外付费的数据
- `404`: 查询的数据或地区不存在
- `429`: 超过访问次数或访问过于频繁
- `500`: 无响应或超时，接口服务异常请联系我们

更多状态码信息请参考[和风天气API文档](https://dev.qweather.com/docs/resource/status-code/)。

## 许可证

MIT License

## 贡献

欢迎提交Issue和Pull Request！

## 相关链接

- [和风天气开发者平台](https://dev.qweather.com/)
- [身份认证文档](https://dev.qweather.com/docs/resource/auth/)
- [JWT认证详细说明](https://dev.qweather.com/docs/resource/auth/#json-web-token)
- [格点实时天气API文档](https://dev.qweather.com/docs/api/weather/grid-weather-now/)
- [状态码说明](https://dev.qweather.com/docs/resource/status-code/)
- [天气图标](https://icons.qweather.com/)
- [控制台-项目管理](https://console.qweather.com/project) 