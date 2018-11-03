package main

import (
	"bufio"
	"log"
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/joeshaw/envdecode"
)

type TwitterUtils struct {
	Keys TwitterKey
	Api  anaconda.TwitterApi
}

type TwitterKey struct {
	ConsumerKey       string `env:"TWITTER_CONSUMER_KEY,required"`
	ConsumerKeySecret string `env:"TWITTER_CONSUMER_SECRET,required"`
	AccessToken       string `env:"TWITTER_ACCESS_TOKEN_KEY,required"`
	AccessTokenSecret string `env:"TWITTER_ACCESS_TOKEN_SECRET,required"`
	ScreenName        string `env:"TWITTER_SCREEN_NAME,required"`
}

var (
	tk TwitterKey
)

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

	var tw TwitterUtils
	tw.Setup()
	for i := 0; i < len(rmsg); i += tweet_max_length {
		if i+tweet_max_length < len(rmsg) {
			tw.PostTweet(string(rmsg[i : i+tweet_max_length]))
		} else {
			tw.PostTweet(string(rmsg[i:]))
		}
	}
}

func (t *TwitterUtils) Setup() {
	// Get Keys of your Environment variables
	if err := envdecode.Decode(&t.Keys); err != nil {
		log.Fatal(err)
	}
	anaconda.SetConsumerKey(t.Keys.ConsumerKey)
	anaconda.SetConsumerSecret(t.Keys.ConsumerKeySecret)
	t.Api = *anaconda.NewTwitterApi(t.Keys.AccessToken, t.Keys.AccessTokenSecret)
	return
}

func (t *TwitterUtils) PostTweet(msg string) {
	v := url.Values{}
	v.Set("screen_name", t.Keys.ScreenName)

	_, errp := t.Api.PostTweet(msg, v)
	if errp != nil {
		log.Fatal(errp)
	}
	return
}
