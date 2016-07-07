package main

import "testing"

func TestConvertToString(t *testing.T) {

	fakeStruct := Forecast{
		Humidity:     82,
		High:         64.99,
		Low:          53.01,
		WindSpeed:    8.05,
		Clouds:       13,
		Weather:      "mist",
		WeatherLevel: "fog",
		WeatherId:    801,
	}
	expected := "weather: :cloud: humidity: 82, high: 64.99, low: 53.01, windspeed: 8.05 "
	result := ConvertToString(fakeStruct)
	if result != expected {
		t.Fatalf("Expected %s, got %s", expected, result)
	}
}
