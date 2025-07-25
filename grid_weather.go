package qweather

import (
	"encoding/json"
	"fmt"
	"net/url"
	"regexp"
)

// GetGridWeatherNow retrieves current grid weather data
// location: longitude,latitude coordinates (e.g., "116.41,39.92")
// options: optional parameters for language and unit settings
func (c *Client) GetGridWeatherNow(location string, options ...GridWeatherNowOptions) (*GridWeatherNowResponse, error) {
	if location == "" {
		return nil, fmt.Errorf("location is required")
	}

	// Validate location format (longitude,latitude)
	if !isValidLocation(location) {
		return nil, fmt.Errorf("invalid location format, expected 'longitude,latitude' (e.g., '116.41,39.92')")
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

	body, err := c.makeRequest("/v7/grid-weather/now", params)
	if err != nil {
		return nil, fmt.Errorf("failed to get grid weather now: %w", err)
	}

	var response GridWeatherNowResponse
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

// isValidLocation validates the location format (longitude,latitude)
func isValidLocation(location string) bool {
	// Regular expression to match longitude,latitude format
	// Longitude: -180 to 180, Latitude: -90 to 90
	// Supports up to 2 decimal places as per API documentation
	pattern := `^-?(?:180(?:\.0{1,2})?|(?:1[0-7]\d|[1-9]?\d)(?:\.\d{1,2})?),-?(?:90(?:\.0{1,2})?|(?:[1-8]?\d)(?:\.\d{1,2})?)$`
	matched, _ := regexp.MatchString(pattern, location)
	return matched
}

// GetGridWeatherNowWithCoordinates is a convenience method that takes separate longitude and latitude
func (c *Client) GetGridWeatherNowWithCoordinates(longitude, latitude float64, options ...GridWeatherNowOptions) (*GridWeatherNowResponse, error) {
	location := fmt.Sprintf("%.2f,%.2f", longitude, latitude)
	return c.GetGridWeatherNow(location, options...)
} 