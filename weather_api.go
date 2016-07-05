package main

import (
  owm "github.com/briandowns/openweathermap"
  "log"
  "os"
)

type Forecast struct {
  Humidity int
  High     float64
  Low      float64
  WindSpeed float64
}

type WeatherApi struct {
  Location string
}

func NewWeatherApi(apiKey string) *WeatherApi {
  os.Setenv("OWM_API_KEY", apiKey)

  return &WeatherApi{
    Location: "San Francisco, CA",
  }
}

func (w *WeatherApi) Now() Forecast {
  weather, err := owm.NewCurrent("F", "EN") // fahrenheit (imperial) with Russian output
  if err != nil {
    log.Fatalln(err)
  }

  weather.CurrentByName(w.Location)

  return Forecast{
    Humidity: weather.Main.Humidity,
    High: weather.Main.TempMax,
    Low: weather.Main.TempMin,
    WindSpeed: weather.Wind.Speed,
  }
}


