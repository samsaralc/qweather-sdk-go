package main

import (
	"fmt"
	"log"
	"testing"

	"github.com/samsaralc/qweather-sdk-go"
)

// TestGridWeatherNow 测试格点实时天气功能
func TestGridWeatherNow(t *testing.T) {
	fmt.Println("=== 格点实时天气测试 ===")

	// 使用API Key认证创建客户端
	client := qweather.NewClientWithAPIKey("your_api_key_here", "your_host_here")

	// 测试1: 使用经纬度字符串获取格点天气
	fmt.Println("\n--- 测试1: 使用位置字符串 ---")
	location := "116.41,39.92" // 北京的经纬度
	response, err := client.GetGridWeatherNow(location)
	if err != nil {
		log.Printf("获取格点天气数据失败: %v", err)
	} else {
		printGridWeatherInfo(response)
	}

	// 测试2: 使用经纬度数值获取格点天气
	fmt.Println("\n--- 测试2: 使用经纬度数值 ---")
	longitude := 121.47
	latitude := 31.23
	response2, err := client.GetGridWeatherNowWithCoordinates(longitude, latitude)
	if err != nil {
		log.Printf("获取格点天气数据失败: %v", err)
	} else {
		printGridWeatherInfo(response2)
	}

	// 测试3: 使用可选参数（语言和单位）
	fmt.Println("\n--- 测试3: 使用可选参数 ---")
	options := qweather.GridWeatherNowOptions{
		Lang: "en", // 英文
		Unit: "i",  // 英制单位
	}
	response3, err := client.GetGridWeatherNow(location, options)
	if err != nil {
		log.Printf("获取格点天气数据失败: %v", err)
	} else {
		printGridWeatherInfo(response3)
	}

	// 测试4: 错误处理 - 无效的坐标格式
	fmt.Println("\n--- 测试4: 错误处理测试 ---")
	invalidLocation := "invalid_coordinates"
	_, err = client.GetGridWeatherNow(invalidLocation)
	if err != nil {
		fmt.Printf("期望的错误: %v\n", err)
	}

	// 测试5: 使用API Key认证
	fmt.Println("\n--- 测试5: API Key认证 ---")
	clientAPIKey := qweather.NewClientWithAPIKey("your_api_key_here", "your_host_here")
	response5, err := clientAPIKey.GetGridWeatherNow(location)
	if err != nil {
		log.Printf("API Key认证失败: %v", err)
	} else {
		fmt.Printf("API Key认证成功，状态码: %s\n", response5.Code)
	}
}

// printGridWeatherInfo 打印格点天气信息的辅助函数
func printGridWeatherInfo(response *qweather.GridWeatherNowResponse) {
	fmt.Printf("状态码: %s\n", response.Code)
	fmt.Printf("更新时间: %s\n", response.UpdateTime.Format("2006-01-02 15:04:05"))
	fmt.Printf("观测时间: %s\n", response.Now.ObsTime.Format("2006-01-02 15:04:05"))
	fmt.Printf("温度: %s°C\n", response.Now.Temp)
	fmt.Printf("天气状况: %s (图标: %s)\n", response.Now.Text, response.Now.Icon)
	fmt.Printf("风向: %s (%s°)\n", response.Now.WindDir, response.Now.Wind360)
	fmt.Printf("风力等级: %s级，风速: %s km/h\n", response.Now.WindScale, response.Now.WindSpeed)
	fmt.Printf("相对湿度: %s%%\n", response.Now.Humidity)
	fmt.Printf("降水量: %s mm\n", response.Now.Precip)
	fmt.Printf("大气压强: %s hPa\n", response.Now.Pressure)

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
