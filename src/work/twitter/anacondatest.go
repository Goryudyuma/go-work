package main

import (
	"bufio"
	"github.com/ChimeraCoder/anaconda"
	"github.com/garyburd/redigo/redis"
	"net/url"
	"os"
)

func main() {
	fp, err := os.Open("/etc/pass/twi.txt")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		panic(err)
	}
	defer c.Close()
	c.Do("SELECT", "2")

	scanner := bufio.NewScanner(fp)

	scanner.Scan()
	anaconda.SetConsumerKey(scanner.Text())
	scanner.Scan()
	anaconda.SetConsumerSecret(scanner.Text())

	scanner.Scan()
	accesstoken := scanner.Text()
	scanner.Scan()
	accesstoken_secret := scanner.Text()
	api := anaconda.NewTwitterApi(accesstoken, accesstoken_secret)
	//api.SetLogger(anaconda.BasicLogger) // logger を設定

	v := url.Values{}
	stream := api.UserStream(v) // 接続

	for {
		// 受信待ち
		select {
		case item := <-stream.C:
			switch status := item.(type) {
			case anaconda.Tweet:
				// Tweet を受信
				//fmt.Printf("%s: %s\n", status.User.ScreenName, status.Text)
				c.Do("LPUSH", status.User.ScreenName, status)
				c.Do("LTRIM", status.User.ScreenName, "1000")
			default:
			}
		}
	}
}
