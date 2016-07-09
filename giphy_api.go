package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	Endpoint = "http://api.giphy.com"
)

//Adapted from github.com/ivolo/go-giphy
type GiphyAPI struct {
	Key string
}

type RandomResponse struct {
	Data Gif `json:"data"`
}

type Gif struct {
	URL string `json:"url"`
}

func NewGiphyAPI(key string) *GiphyAPI {
	c := &GiphyAPI{Key: key}
	return c
}

func (c *GiphyAPI) Random(query string) (*Gif, error) {
	u := fmt.Sprintf("%s/v1/gifs/random?tag=%s&api_key=%s", Endpoint, url.QueryEscape(query), c.Key)
	res, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var s RandomResponse
	// fmt.Println(string(body))
	if err := json.Unmarshal(body, &s); err != nil {
		return nil, err
	}
	return &s.Data, nil
}
