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

	forecast := api.Now()

	emojis := ConvertToEmoji(ConvertToString(forecast))

	twitterAPI.PostTweet(emojis)

	fmt.Println(twitterAPI)

	hour := (time.Now().Hour())
	if hour == 9 || hour == 14 || hour == 20 {
		fmt.Println(emojis)
	}
}
