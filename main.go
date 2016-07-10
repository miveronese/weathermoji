package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	api := NewWeatherApi(os.Getenv("OPENWEATHER_API_KEY"))
	twitterAPI := NewTwitterService(
		os.Getenv("TWITTER_CONSUMER_KEY"),
		os.Getenv("TWITTER_CONSUMER_SECRET"),
		os.Getenv("TWITTER_API_TOKEN"),
		os.Getenv("TWITTER_API_SECRET"),
	)
	giphyClient := NewGiphyAPI(os.Getenv("GIPHY_API_KEY"))

	// urlGif := GrabGiphy("#" + forecast.WeatherLevel + " #weather")

	for {
		forecast := api.Now()
		emojis := ConvertToString(forecast)
		fmt.Println(time.Now())
		fmt.Println(forecast.Weather)
		gif, _ := giphyClient.Random(forecast.WeatherLevel)

		fmt.Println(emojis + " " + gif.URL)
		fmt.Println(gif.URL)
		twitterAPI.PostTweet(emojis + " " + gif.URL)
		time.Sleep(2 * time.Hour)
		fmt.Println(twitterAPI)
	}
}
