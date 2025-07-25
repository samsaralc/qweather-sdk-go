package qweather

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Client represents the QWeather API client
type Client struct {
	BaseURL    string
	Token      string
	AuthType   AuthType
	HTTPClient *http.Client
}

// AuthType represents the authentication type
type AuthType int

const (
	AuthTypeJWT    AuthType = iota // JWT Bearer Token (推荐)
	AuthTypeAPIKey                 // API KEY (传统方式)
)

// Config holds the configuration for the QWeather client
type Config struct {
	Host     string        // API基础URL，默认: "https://devapi.qweather.com"
	Token    string        // JWT Token 或 API Key
	AuthType AuthType      // 认证方式：JWT 或 API Key，默认为JWT
	Timeout  time.Duration // HTTP请求超时时间，默认: 30秒
}

// NewClient creates a new QWeather API client
func NewClient(config Config) *Client {
	if config.Host == "" {
		config.Host = "https://devapi.qweather.com"
	}

	// 确保Host包含协议前缀
	if config.Host != "" && !hasProtocolScheme(config.Host) {
		config.Host = "https://" + config.Host
	}

	if config.Timeout == 0 {
		config.Timeout = 30 * time.Second
	}

	// 默认使用JWT认证方式
	authType := config.AuthType
	if authType != AuthTypeJWT && authType != AuthTypeAPIKey {
		authType = AuthTypeJWT
	}

	return &Client{
		BaseURL:  config.Host,
		Token:    config.Token,
		AuthType: authType,
		HTTPClient: &http.Client{
			Timeout: config.Timeout,
		},
	}
}

// makeRequest performs HTTP request with authentication
func (c *Client) makeRequest(endpoint string, params url.Values) ([]byte, error) {
	if c.Token == "" {
		return nil, fmt.Errorf("API token is required")
	}

	u, err := url.Parse(c.BaseURL + endpoint)
	if err != nil {
		return nil, fmt.Errorf("invalid URL: %w", err)
	}

	u.RawQuery = params.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// 根据认证类型设置不同的认证头
	switch c.AuthType {
	case AuthTypeJWT:
		req.Header.Set("Authorization", "Bearer "+c.Token)
	case AuthTypeAPIKey:
		req.Header.Set("X-QW-Api-Key", c.Token)
	default:
		// 默认使用JWT认证
		req.Header.Set("Authorization", "Bearer "+c.Token)
	}

	req.Header.Set("Accept-Encoding", "gzip")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	var reader io.Reader = resp.Body

	// 处理GZIP压缩响应
	if strings.Contains(resp.Header.Get("Content-Encoding"), "gzip") {
		gzipReader, err := gzip.NewReader(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to create gzip reader: %w", err)
		}
		defer gzipReader.Close()
		reader = gzipReader
	}

	body, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	return body, nil
}

// hasProtocolScheme 检查URL是否已经包含协议前缀
func hasProtocolScheme(u string) bool {
	return len(u) > 7 && (u[:7] == "http://" || u[:8] == "https://")
}

// NewClientWithJWT creates a new QWeather API client with JWT authentication
func NewClientWithJWT(jwtToken string, host string) *Client {
	return NewClient(Config{
		Host:     host,
		Token:    jwtToken,
		AuthType: AuthTypeJWT,
	})
}

// NewClientWithAPIKey creates a new QWeather API client with API Key authentication
func NewClientWithAPIKey(apiKey string, host string) *Client {
	return NewClient(Config{
		Host:     host,
		Token:    apiKey,
		AuthType: AuthTypeAPIKey,
	})
}
