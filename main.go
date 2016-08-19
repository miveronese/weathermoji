package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/ChimeraCoder/anaconda"
)

func updateForecast(api WeatherApi, twitterAPI TwitterService, giphyClient *GiphyAPI, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		forecast := api.Now()
		emojis := ConvertToString(forecast)
		gif, _ := giphyClient.Random(forecast.WeatherLevel)
		fmt.Println(emojis + " " + gif.URL)
		// twitterAPI.PostTweet(emojis + " " + gif.URL)
		fmt.Println(twitterAPI)
		time.Sleep(5 * time.Second)
	}
}

func replyTweets(mentions []anaconda.Tweet, wg *sync.WaitGroup) {
	defer wg.Done()
	for x := range mentions {
		p := fmt.Println
		p(mentions[x].User)
		p(mentions[x].Retweeted)
		p(mentions[x].Text)
		p(mentions[x].Id)
		p(mentions[x].InReplyToStatusIdStr)
		p(mentions[x].InReplyToUserID)
		p(mentions[x].InReplyToScreenName)
		//result, _ := time.Parse("Mon Jan _2 15:04:05 +0000 2006", mentions[x].CreatedAt)
		//p(result.Format("2006-01-02"))
		//p(time.Now().Format("2006-01-02"))
	}
}

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
	// fmt.Println(convertToEmoji(":cloud:"))
	mentions, _ := twitterAPI.GetMentionsTimeline()

	var wg sync.WaitGroup
	wg.Add(2)
	go updateForecast(api, twitterAPI, giphyClient, &wg)
	go replyTweets(mentions, &wg)
	wg.Wait()
}
