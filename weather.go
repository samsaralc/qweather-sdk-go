package qweather

import (
	"encoding/json"
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

// GetWeatherNow retrieves current weather data for cities
// location: LocationID (e.g., "101010100") or longitude,latitude coordinates (e.g., "116.41,39.92")
// options: optional parameters for language and unit settings
func (c *Client) GetWeatherNow(location string, options ...WeatherNowOptions) (*WeatherNowResponse, error) {
	if location == "" {
		return nil, fmt.Errorf("location is required")
	}

	// Validate location format (LocationID or longitude,latitude)
	if !isValidLocationIDOrCoordinate(location) {
		return nil, fmt.Errorf("invalid location format, expected LocationID (e.g., '101010100') or 'longitude,latitude' (e.g., '116.41,39.92')")
	}

	params := url.Values{}
	params.Set("location", location)

	// Apply optional parameters
	if len(options) > 0 {
		opt := options[0]
		if opt.Lang != "" {
			params.Set("lang", opt.Lang)
		}
		if opt.Unit != "" {
			params.Set("unit", opt.Unit)
		}
	}

	body, err := c.makeRequest("/v7/weather/now", params)
	if err != nil {
		return nil, fmt.Errorf("failed to get weather now: %w", err)
	}

	var response WeatherNowResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	// Check API response code
	if response.Code != "200" {
		return nil, Error{
			Code:    response.Code,
			Message: fmt.Sprintf("API returned error code: %s", response.Code),
		}
	}

	return &response, nil
}

// isValidLocationIDOrCoordinate validates the location format (LocationID or longitude,latitude)
func isValidLocationIDOrCoordinate(location string) bool {
	// Check if it's a LocationID (numeric, typically 9 digits)
	if matched, _ := regexp.MatchString(`^\d{6,12}$`, location); matched {
		return true
	}

	// Check if it's coordinates (longitude,latitude)
	if strings.Contains(location, ",") {
		return isValidLocation(location)
	}

	return false
}

// GetWeatherNowWithLocationID is a convenience method for using LocationID
func (c *Client) GetWeatherNowWithLocationID(locationID string, options ...WeatherNowOptions) (*WeatherNowResponse, error) {
	// Validate LocationID format
	if matched, _ := regexp.MatchString(`^\d{6,12}$`, locationID); !matched {
		return nil, fmt.Errorf("invalid LocationID format, expected numeric string (e.g., '101010100')")
	}

	return c.GetWeatherNow(locationID, options...)
}

// GetWeatherNowWithCoordinates is a convenience method for using coordinates
func (c *Client) GetWeatherNowWithCoordinates(longitude, latitude float64, options ...WeatherNowOptions) (*WeatherNowResponse, error) {
	location := fmt.Sprintf("%.2f,%.2f", longitude, latitude)
	return c.GetWeatherNow(location, options...)
}
