package weather_test

import (
	"http/geo"
	"http/weather"
	"strings"
	"testing"
)

func TestGetWeather(t *testing.T) {
	excpected := "London"
	geoData := geo.GeoData{
		City: excpected,
	}
	format := 3

	result := weather.GetWeather(geoData, format)
	if !strings.Contains(result, excpected){
		t.Errorf("Ожидалось %v, получено %v", excpected, result)
	}
}