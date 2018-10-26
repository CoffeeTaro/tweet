# tweet

標準入力につぶやく内容を与えると,Twitterにその内容を投稿するプログラム


### Download

```
$ git clone http://github.com/kotaru23/tweet
```

### 環境変数の設定
.bashrcまたは.bash_profileに下記を記述  
.bash_profileに書くことを推奨  

```
# Twitter API Keys
export TWITTER_CONSUMER_KEY="your_twiter_consumer_key"
export TWITTER_CONSUMER_SECRET="your_consumer_secret"
export TWITTER_ACCESS_TOKEN_KEY="your_access_token_key"
export TWITTER_ACCESS_TOKEN_SECRET="your_access_token_secret"
# Twitter Screen Name
export TWITTER_SCREEN_NAME="your_screen_name"
```

## 実行例

```
$ echo hoge | tweet
```
