package main

import (
	"fmt"

	"github.com/ChimeraCoder/anaconda"
)

type TwitterService struct {
	Api *anaconda.TwitterApi
}

func NewTwitterService(consumerKey string, consumerSecret string, apiToken string, apiSecret string) TwitterService {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(apiToken, apiSecret)

	return TwitterService{
		Api: api,
	}
}

func (self TwitterService) PostTweet(text string) {
	self.Api.PostTweet(text, nil)
}

func (self TwitterService) GetSearch(text string) {
	searchResult, err := self.Api.GetSearch(text, nil)
	if err != nil {
		panic(err)
	}
	for _, tweet := range searchResult.Statuses {
		fmt.Print(tweet.Text)
	}
}

func (self TwitterService) GetMentionsTimeline() (timeline []anaconda.Tweet, err error) {
	return self.Api.GetMentionsTimeline(nil)
}

// Not being used yet
func (self TwitterService) GetTweet(id int64) (timeline anaconda.Tweet, err error) {
	return self.Api.GetTweet(id, nil)
}
