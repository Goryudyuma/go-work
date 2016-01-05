package main

import (
	"bufio"
	"github.com/ChimeraCoder/anaconda"
	"os"
	"runtime"
)

//globalフラグ
var r, h *bool

//認証関数
func access(fp *os.File) *anaconda.TwitterApi {
	scanner := bufio.NewScanner(fp)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	scanner.Scan()
	consumerKey := scanner.Text()
	scanner.Scan()
	consumerSecret := scanner.Text()
	scanner.Scan()
	accesstoken := scanner.Text()
	scanner.Scan()
	accesstokenSecret := scanner.Text()

	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	return anaconda.NewTwitterApi(accesstoken, accesstokenSecret)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	apiChan := make(chan *anaconda.TwitterApi)

	go func() {
		fp, err := os.Open("/etc/pass/twi.txt")
		if err != nil {
			panic(err)
		}
		defer fp.Close()

		apiChan <- access(fp)
	}()

	api := <-apiChan

	api.Retweet(684185068345016320, true)
}
