# tweet

Simple Command Line Twitter Client written in Golang.  
Read stdin charactors, then post them to Twitter.

### Install

```
$ go get github.com/kotaru23/tweet
```

### Set Environment Variables

Write these on your ~/.bash_profile

```
# Twitter API Keys
export TWITTER_CONSUMER_KEY="your_twiter_consumer_key"
export TWITTER_CONSUMER_SECRET="your_consumer_secret"
export TWITTER_ACCESS_TOKEN_KEY="your_access_token_key"
export TWITTER_ACCESS_TOKEN_SECRET="your_access_token_secret"
# Twitter Screen Name
export TWITTER_SCREEN_NAME="your_screen_name"
```

## Example

```
$ echo "I'm hungry." | tweet
```
