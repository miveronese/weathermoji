package main

import (
	"fmt"
	"os"
)

func main() {

	api := NewWeatherApi(os.Getenv("OPENWEATHER_API_KEY"))

	twitterApi := NewTwitterService(
		os.Getenv("TWITTER_CONSUMER_KEY"),
		os.Getenv("TWITTER_CONSUMER_SECRET"),
		os.Getenv("TWITTER_API_TOKEN"),
		os.Getenv("TWITTER_API_SECRET"),
	)

	forecast := api.Now()

	emojis := ConvertToEmoji(ConvertToString(forecast))

	twitterApi.PostTweet(emojis)

	fmt.Println(twitterApi)
	fmt.Println(emojis)
}
