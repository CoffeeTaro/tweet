#!/bin/bash

docker run --rm -i \
    -e TWITTER_CONSUMER_KEY=$TWITTER_CONSUMER_KEY \
    -e TWITTER_CONSUMER_SECRET=$TWITTER_CONSUMER_SECRET \
    -e TWITTER_ACCESS_TOKEN_KEY=$TWITTER_ACCESS_TOKEN_KEY \
    -e TWITTER_ACCESS_TOKEN_SECRET=$TWITTER_ACCESS_TOKEN_SECRET \
    -e TWITTER_SCREEN_NAME=$TWITTER_SCREEN_NAME \
    kotaru/tweet
