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
	expected := "☁️   🕛 0˚F  ⬆️ 64.99˚F  ⬇️ 53.01˚F #fog #weather #sanfrancisco"
	result := ConvertToString(fakeStruct)
	if result != expected {
		t.Fatalf("Expected %s, got %s", expected, result)
	}
}
