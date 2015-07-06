package main

import (
	"bufio"
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"github.com/otiai10/twistream"
	"os"
)

func main() {
	var fp *os.File
	var err error
	fp, err = os.Open("/etc/pass/twi.txt")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	scanner.Scan()
	consumer_key := scanner.Text()
	scanner.Scan()
	consumer_secret := scanner.Text()
	scanner.Scan()
	accesstoken := scanner.Text()
	scanner.Scan()
	accesstoken_secret := scanner.Text()

	anaconda.SetConsumerKey(consumer_key)
	anaconda.SetConsumerSecret(consumer_secret)
	//api := anaconda.NewTwitterApi(accesstoken, accesstoken_secret)

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	//	tweet, err := api.PostTweet("test tweet2", nil)
	//	fmt.Println(tweet)

	timeline, _ := twistream.New(
		"https://userstream.twitter.com/1.1/user.json",
		consumer_key,
		consumer_secret,
		accesstoken,
		accesstoken_secret,
	)

	for {
		status := <-timeline.Listen()
		fmt.Println(status.User.Name + "  @" + status.User.ScreenName)
		fmt.Println(status.Text)
		fmt.Println()
	}
}
