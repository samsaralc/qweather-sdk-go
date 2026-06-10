package main

import (
	"fmt"
	"log"
	"testing"

	"github.com/samsaralc/qweather-sdk-go"
)

// TestWeatherNow 测试实时天气功能
func TestWeatherNow(t *testing.T) {
	fmt.Println("=== 实时天气测试 ===")

	// 使用API Key认证创建客户端
	client := qweather.NewClientWithAPIKey("your_api_key_here", "your_host_here")

	// 测试1: 使用LocationID获取实时天气
	fmt.Println("\n--- 测试1: 使用LocationID ---")
	locationID := "101010100" // 北京的LocationID
	response, err := client.GetWeatherNow(locationID)
	if err != nil {
		log.Printf("获取实时天气数据失败: %v", err)
	} else {
		printCityWeatherInfo(response)
	}

	// 测试2: 使用LocationID便利方法
	fmt.Println("\n--- 测试2: 使用LocationID便利方法 ---")
	shanghaiID := "101020100" // 上海的LocationID
	response2, err := client.GetWeatherNowWithLocationID(shanghaiID)
	if err != nil {
		log.Printf("获取实时天气数据失败: %v", err)
	} else {
		printCityWeatherInfo(response2)
	}

	// 测试3: 使用经纬度坐标获取实时天气
	fmt.Println("\n--- 测试3: 使用经纬度坐标 ---")
	coordinates := "116.41,39.92" // 北京的经纬度
	response3, err := client.GetWeatherNow(coordinates)
	if err != nil {
		log.Printf("获取实时天气数据失败: %v", err)
	} else {
		printCityWeatherInfo(response3)
	}

	// 测试4: 使用经纬度数值获取实时天气
	fmt.Println("\n--- 测试4: 使用经纬度数值 ---")
	longitude := 121.47
	latitude := 31.23
	response4, err := client.GetWeatherNowWithCoordinates(longitude, latitude)
	if err != nil {
		log.Printf("获取实时天气数据失败: %v", err)
	} else {
		printCityWeatherInfo(response4)
	}

	// 测试5: 使用可选参数（语言和单位）
	fmt.Println("\n--- 测试5: 使用可选参数 ---")
	options := qweather.WeatherNowOptions{
		Lang: "en", // 英文
		Unit: "i",  // 英制单位
	}
	response5, err := client.GetWeatherNow(locationID, options)
	if err != nil {
		log.Printf("获取实时天气数据失败: %v", err)
	} else {
		printCityWeatherInfo(response5)
	}

	// 测试6: 错误处理 - 无效的LocationID格式
	fmt.Println("\n--- 测试6: 错误处理测试 ---")
	invalidLocationID := "invalid_id"
	_, err = client.GetWeatherNowWithLocationID(invalidLocationID)
	if err != nil {
		fmt.Printf("期望的错误: %v\n", err)
	}

	// 测试7: 错误处理 - 无效的位置格式
	fmt.Println("\n--- 测试7: 无效位置格式测试 ---")
	invalidLocation := "invalid_location"
	_, err = client.GetWeatherNow(invalidLocation)
	if err != nil {
		fmt.Printf("期望的错误: %v\n", err)
	}

	// 测试8: 使用API Key认证
	fmt.Println("\n--- 测试8: API Key认证 ---")
	clientAPIKey := qweather.NewClientWithAPIKey("your_api_key_here", "your_host_here")
	response8, err := clientAPIKey.GetWeatherNow(locationID)
	if err != nil {
		log.Printf("API Key认证失败: %v", err)
	} else {
		fmt.Printf("API Key认证成功，状态码: %s\n", response8.Code)
	}
}

// printCityWeatherInfo 打印实时天气信息的辅助函数
func printCityWeatherInfo(response *qweather.WeatherNowResponse) {
	fmt.Printf("状态码: %s\n", response.Code)
	fmt.Printf("更新时间: %s\n", response.UpdateTime.Format("2006-01-02 15:04:05"))
	fmt.Printf("观测时间: %s\n", response.Now.ObsTime.Format("2006-01-02 15:04:05"))
	fmt.Printf("温度: %s°C\n", response.Now.Temp)

	if response.Now.FeelsLike != "" {
		fmt.Printf("体感温度: %s°C\n", response.Now.FeelsLike)
	}

	fmt.Printf("天气状况: %s (图标: %s)\n", response.Now.Text, response.Now.Icon)
	fmt.Printf("风向: %s (%s°)\n", response.Now.WindDir, response.Now.Wind360)
	fmt.Printf("风力等级: %s级，风速: %s km/h\n", response.Now.WindScale, response.Now.WindSpeed)
	fmt.Printf("相对湿度: %s%%\n", response.Now.Humidity)
	fmt.Printf("降水量: %s mm\n", response.Now.Precip)
	fmt.Printf("大气压强: %s hPa\n", response.Now.Pressure)

	if response.Now.Vis != "" {
		fmt.Printf("能见度: %s km\n", response.Now.Vis)
	}
	if response.Now.Cloud != "" {
		fmt.Printf("云量: %s%%\n", response.Now.Cloud)
	}
	if response.Now.Dew != "" {
		fmt.Printf("露点温度: %s°C\n", response.Now.Dew)
	}

	fmt.Printf("响应式页面: %s\n", response.FxLink)

	if len(response.Refer.Sources) > 0 {
		fmt.Printf("数据源: %v\n", response.Refer.Sources)
	}
	if len(response.Refer.License) > 0 {
		fmt.Printf("许可证: %v\n", response.Refer.License)
	}
}
