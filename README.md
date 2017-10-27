# TweetCommand

標準入力につぶやく内容を与えると,Twitterにその内容を投稿するスクリプト  
開発中  

# 環境
macOS, Ubuntu, CentOSなど，シェルが使えるOS
Go言語で記述されているので，Windowsでもコンパイルしてやれば動くはず(未確認)


## 準備

### ダウンロード
必要があればコンパイルする  

```
# ダウンロード
$ git clone http://github.com/CoffeeTaro/tweet
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

Cronなどで，指定した時刻にTweetコマンドを実行したい場合は/etc/environmentに下記のように記述してもいいでしょう

```
# Twitter API Keys
TWITTER_CONSUMER_KEY="your_twiter_consumer_key 
TWITTER_CONSUMER_SECRET="your_consumer_secret"
TWITTER_ACCESS_TOKEN_KEY="your_access_token_key"
TWITTER_ACCESS_TOKEN_SECRET="your_access_token_secret"
# Twitter Screen Name
TWITTER_SCREEN_NAME="your_screen_name"
```


## 実行例

```
$ echo hoge | tweet
```
