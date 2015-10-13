package main

import (
	"bufio"
	"encoding/json"
	"github.com/ChimeraCoder/anaconda"
	"github.com/garyburd/redigo/redis"
	"net/url"
	"os"
	"strings"
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
	c.Do("SELECT", "3")

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
				b, err := json.Marshal(status)
				if err != nil {
					continue
				}
				c.Do("SETEX", status.Id, 86400, b)
			case anaconda.StatusDeletionNotice:
				val, err := redis.String(c.Do("GET", status.IdStr))
				if err == nil {
					obj := new(anaconda.Tweet)
					dec := json.NewDecoder(strings.NewReader(val))
					dec.Decode(&obj)
					c.Do("LPUSH", strings.ToLower(obj.User.ScreenName), val)
					c.Do("LTRIM", strings.ToLower(obj.User.ScreenName), "0", "999")
					c.Do("LPUSH", "all", val)
					c.Do("LTRIM", "all", "0", "999")
					c.Do("DEL", status.IdStr)
				}
			case anaconda.DirectMessage:
				if status.SenderId == 119667108 {
					classification := strings.Fields(status.Text)
					if classification[0] == "del" {
						val, err := redis.String(c.Do("LINDEX", strings.ToLower(classification[1]), classification[2]))
						if err == nil {
							api.PostDMToUserId(val+"\n"+status.Text, 119667108)
						} else {
							api.PostDMToUserId("Nothing."+"\n"+status.Text, 119667108)
						}
					}
				}
			default:
				//fmt.Println(status)
			}
		}
	}
}
