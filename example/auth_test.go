package main

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/samsaralc/qweather-sdk-go"
)

// TestAuthentication 测试身份认证功能
func TestAuthentication(t *testing.T) {
	fmt.Println("=== 身份认证测试 ===")

	// 测试1: JWT认证
	fmt.Println("\n--- 测试1: JWT认证 ---")
	testJWTAuthentication()

	// 测试2: API Key认证
	fmt.Println("\n--- 测试2: API Key认证 ---")
	testAPIKeyAuthentication()

	// 测试3: 完整配置方式
	fmt.Println("\n--- 测试3: 完整配置方式 ---")
	testFullConfiguration()

	// 测试4: 认证方式对比
	fmt.Println("\n--- 测试4: 认证方式对比 ---")
	testAuthenticationComparison()

	// 测试5: 错误处理测试
	fmt.Println("\n--- 测试5: 错误处理测试 ---")
	testAuthenticationErrors()
}

// testJWTAuthentication 测试JWT认证
func testJWTAuthentication() {
	fmt.Println("创建JWT认证客户端...")

	// 使用便利构造函数创建JWT客户端
	client := qweather.NewClientWithJWT("your_jwt_token_here", "your_host_here")

	fmt.Printf("认证类型: %v\n", client.AuthType)
	fmt.Printf("Token: %s\n", maskToken(client.Token))
	fmt.Printf("Host: %s\n", client.BaseURL)
	fmt.Printf("超时时间: %v\n", client.HTTPClient.Timeout)

	// 测试API调用
	location := "116.41,39.92"
	response, err := client.GetGridWeatherNow(location)
	if err != nil {
		log.Printf("JWT认证API调用失败: %v", err)
	} else {
		fmt.Printf("JWT认证成功，状态码: %s\n", response.Code)
	}
}

// testAPIKeyAuthentication 测试API Key认证
func testAPIKeyAuthentication() {
	fmt.Println("创建API Key认证客户端...")

	// 使用便利构造函数创建API Key客户端
	client := qweather.NewClientWithAPIKey("your_api_key_here", "your_api_key_host_here")

	fmt.Printf("认证类型: %v\n", client.AuthType)
	fmt.Printf("Token: %s\n", maskToken(client.Token))
	fmt.Printf("Host: %s\n", client.BaseURL)
	fmt.Printf("超时时间: %v\n", client.HTTPClient.Timeout)

	// 测试API调用
	locationID := "101010100"
	response, err := client.GetWeatherNow(locationID)
	if err != nil {
		log.Printf("API Key认证API调用失败: %v", err)
	} else {
		fmt.Printf("API Key认证成功，状态码: %s\n", response.Code)
	}
}

// testFullConfiguration 测试完整配置方式
func testFullConfiguration() {
	fmt.Println("使用完整配置创建客户端...")

	// JWT配置
	jwtConfig := qweather.Config{
		Host:     "https://devapi.qweather.com",
		Token:    "your_jwt_token_here",
		AuthType: qweather.AuthTypeJWT,
		Timeout:  45 * time.Second,
	}

	jwtClient := qweather.NewClient(jwtConfig)
	fmt.Printf("JWT客户端 - 认证类型: %v, 超时: %v\n",
		jwtClient.AuthType, jwtClient.HTTPClient.Timeout)

	// API Key配置
	apiKeyConfig := qweather.Config{
		Host:     "https://devapi.qweather.com",
		Token:    "your_api_key_here",
		AuthType: qweather.AuthTypeAPIKey,
		Timeout:  60 * time.Second,
	}

	apiKeyClient := qweather.NewClient(apiKeyConfig)
	fmt.Printf("API Key客户端 - 认证类型: %v, 超时: %v\n",
		apiKeyClient.AuthType, apiKeyClient.HTTPClient.Timeout)

	// 测试默认配置
	defaultConfig := qweather.Config{
		Token: "default_token_here",
		// AuthType未设置，应该默认为JWT
	}

	defaultClient := qweather.NewClient(defaultConfig)
	fmt.Printf("默认客户端 - 认证类型: %v (应该是JWT), Host: %s, 超时: %v\n",
		defaultClient.AuthType, defaultClient.BaseURL, defaultClient.HTTPClient.Timeout)
}

// testAuthenticationComparison 测试认证方式对比
func testAuthenticationComparison() {
	fmt.Println("认证方式对比测试...")

	location := "116.41,39.92"

	// JWT认证测试
	jwtClient := qweather.NewClientWithJWT("your_jwt_token_here", "your_host_here")
	fmt.Println("JWT认证测试:")
	response1, err1 := jwtClient.GetGridWeatherNow(location)
	if err1 != nil {
		fmt.Printf("  JWT认证失败: %v\n", err1)
	} else {
		fmt.Printf("  JWT认证成功: 状态码 %s\n", response1.Code)
	}

	// API Key认证测试
	apiClient := qweather.NewClientWithAPIKey("your_api_key_here", "your_api_key_host_here")
	fmt.Println("API Key认证测试:")
	response2, err2 := apiClient.GetGridWeatherNow(location)
	if err2 != nil {
		fmt.Printf("  API Key认证失败: %v\n", err2)
	} else {
		fmt.Printf("  API Key认证成功: 状态码 %s\n", response2.Code)
	}

	// 性能对比（模拟）
	fmt.Println("认证方式特点:")
	fmt.Println("  JWT认证: 更安全，支持过期时间，推荐用于生产环境")
	fmt.Println("  API Key认证: 操作简单，但安全性较低，2027年后将限制使用")
}

// testAuthenticationErrors 测试认证错误处理
func testAuthenticationErrors() {
	fmt.Println("认证错误处理测试...")

	// 测试空Token
	fmt.Println("测试空Token:")
	emptyConfig := qweather.Config{
		Host:     "https://devapi.qweather.com",
		Token:    "", // 空Token
		AuthType: qweather.AuthTypeJWT,
	}

	emptyClient := qweather.NewClient(emptyConfig)
	_, err := emptyClient.GetGridWeatherNow("116.41,39.92")
	if err != nil {
		fmt.Printf("  期望的错误: %v\n", err)
	}

	// 测试无效的认证类型设置
	fmt.Println("测试无效认证类型:")
	invalidConfig := qweather.Config{
		Token:    "test_token",
		AuthType: qweather.AuthType(999), // 无效的认证类型
	}

	invalidClient := qweather.NewClient(invalidConfig)
	fmt.Printf("  无效认证类型被重置为: %v (应该是JWT)\n", invalidClient.AuthType)

	// 测试HTTP请求头设置
	fmt.Println("认证头设置测试:")
	fmt.Println("  JWT认证使用: Authorization: Bearer <token>")
	fmt.Println("  API Key认证使用: X-QW-Api-Key: <key>")
}

// maskToken 遮盖Token的敏感部分
func maskToken(token string) string {
	if len(token) <= 8 {
		return "****"
	}
	return token[:4] + "****" + token[len(token)-4:]
}
