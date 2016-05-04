package main

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
)

//認証関数
func access() *anaconda.TwitterApi {
	consumerKey := ""
	consumerSecret := ""
	accesstoken := ""
	accesstokenSecret := ""

	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	return anaconda.NewTwitterApi(accesstoken, accesstokenSecret)
}

func main() {
	tweetid := int64(0)

	apiChan := make(chan *anaconda.TwitterApi)

	go func() {
		apiChan <- access()
	}()

	api := <-apiChan

	tweet, err := api.DeleteTweet(tweetid, true)
	fmt.Println(tweet)
	fmt.Println(err)
}
