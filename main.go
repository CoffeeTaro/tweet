package main

import (
	"bufio"
	"log"
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/joeshaw/envdecode"
)

func post(msg string) {
	var ts struct {
		ConsumerKey       string `env:"TWITTER_CONSUMER_KEY,required"`
		ConsumerKeySecret string `env:"TWITTER_CONSUMER_SECRET,required"`
		AccessToken       string `env:"TWITTER_ACCESS_TOKEN_KEY,required"`
		AccessTokenSecret string `env:"TWITTER_ACCESS_TOKEN_SECRET,required"`
		ScreenName        string `env:"TWITTER_SCREEN_NAME,required"`
	}
	// Get Keys of your Environment variables
	if err := envdecode.Decode(&ts); err != nil {
		log.Fatal(err)
	}
	anaconda.SetConsumerKey(ts.ConsumerKey)
	anaconda.SetConsumerSecret(ts.ConsumerKeySecret)
	api := anaconda.NewTwitterApi(ts.AccessToken, ts.AccessTokenSecret)

	v := url.Values{}
	v.Set("screen_name", ts.ScreenName)

	_, errp := api.PostTweet(msg, v)
	if errp != nil {
		log.Fatal(errp)
	}
}

func main() {
	// read lines
	s := bufio.NewScanner(os.Stdin)
	msg := ""
	for s.Scan() {
		msg = msg + s.Text() + "\n"
		if err := s.Err(); err != nil {
			break
		}
	}
	//convert msg to rune
	rmsg := []rune(msg)
	// post message to twitter
	tweet_max_length := 140
	for i := 0; i < len(rmsg); i += tweet_max_length {
		if i+tweet_max_length < len(rmsg) {
			post(string(rmsg[i : i+tweet_max_length]))
		} else {
			post(string(rmsg[i:]))
		}
	}
}
