package main

import "fmt"

type ConvertInTweet struct {
	Forecast
}

// Forecast format = {84 87.8 57.2 9.17 40 Mist mist 701}

func ConvertToString(forecast Forecast) string {

	return fmt.Sprint("[humidity: ", forecast.Humidity, ", high: ", forecast.High, ", low: ", forecast.Low, " , windspeed: ", forecast.WindSpeed, ", clouds: ", forecast.Clouds, ", weather: ", forecast.Weather, ",  weatherlevel: ", forecast.WeatherLevel, ", weatherid: ", forecast.WeatherId, "]")

}
