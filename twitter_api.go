package main

import (
  "github.com/ChimeraCoder/anaconda"
)

type TwitterService struct {
  Api *anaconda.TwitterApi
}

func NewTwiiterService(consumerKey string, consumerSecret string, apiToken string, apiSecret string) TwitterService { 
  anaconda.SetConsumerKey(consumerKey)
  anaconda.SetConsumerSecret(consumerSecret)
  api := anaconda.NewTwitterApi(apiToken, apiSecret)

  return TwitterService {
    Api: api,
  }
}
