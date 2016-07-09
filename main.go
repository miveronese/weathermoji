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

	forecast := api.Now()
	emojis := ConvertToEmoji(ConvertToString(forecast))
	// urlGif := GrabGiphy("#" + forecast.WeatherLevel + " #weather")

	fmt.Println(twitterAPI)

	for {
		fmt.Println(time.Now())
		fmt.Println(forecast.Weather)
		gif, _ := giphyClient.Random(forecast.WeatherLevel)
		fmt.Println(emojis + " " + gif.URL)

		fmt.Println(gif.URL)
		twitterAPI.PostTweet(emojis + " " + gif.URL)
		time.Sleep(4 * time.Hour)
	}
}
