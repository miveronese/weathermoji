package main

import (
	owm "github.com/briandowns/openweathermap"
	"log"
	"os"
)

type Forecast struct {
	Humidity     int
	High         float64
	Low          float64
	WindSpeed    float64
	Clouds       int
	Weather      string
	WeatherLevel string // TODO map from openweathermap.org/weather-conditions to some constants?
	WeatherId    int
}

type WeatherApi struct {
	Location string
}

func NewWeatherApi(apiKey string) WeatherApi {
	os.Setenv("OWM_API_KEY", apiKey) // owm requires this env var to function

	return WeatherApi{
		Location: "San Francisco, CA",
	}
}

func (self WeatherApi) Now() Forecast {
	weather, err := owm.NewCurrent("F", "EN")
	if err != nil {
		log.Fatalln(err)
	}

	weather.CurrentByName(self.Location)

	return Forecast{
		Humidity:     weather.Main.Humidity,
		High:         weather.Main.TempMax,
		Low:          weather.Main.TempMin,
		WindSpeed:    weather.Wind.Speed,
		Clouds:       weather.Clouds.All,
		Weather:      weather.Weather[0].Main,
		WeatherLevel: weather.Weather[0].Description,
		WeatherId:    weather.Weather[0].ID,
	}
}
