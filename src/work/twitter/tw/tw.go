package main

import (
	"bufio"
	"github.com/ChimeraCoder/anaconda"
	"os"
	"strings"
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
	api := anaconda.NewTwitterApi(accesstoken, accesstoken_secret)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	api.PostTweet(strings.Join(os.Args[1:], " "), nil)
}
