package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"github.com/ChimeraCoder/anaconda"
	"github.com/garyburd/redigo/redis"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func DMTweet(str string, api *anaconda.TwitterApi) {
	api.PostDMToUserId(str, 119667108)
}

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
					switch classification[0] {
					case "del":
						{
							switch len(classification) {
							case 3:
								{
									val, err := redis.String(c.Do("LINDEX", strings.ToLower(classification[1]), classification[2]))
									if err == nil {
										var out bytes.Buffer
										json.Indent(&out, []byte(val), "", "__")
										go DMTweet(out.String()+"\n"+status.Text, api)
									} else {
										go DMTweet("Nothing."+"\n"+status.Text, api)
									}
								}
							case 4:
								{
									b, be := strconv.Atoi(classification[2])
									e, ee := strconv.Atoi(classification[3])
									if be == nil && ee == nil {
										for i := b; i < e; i++ {
											val, err := redis.String(c.Do("LINDEX", strings.ToLower(classification[1]), i))
											if err == nil {
												var out bytes.Buffer
												json.Indent(&out, []byte(val), "", "__")
												go DMTweet(out.String()+"\n"+status.Text+" No."+strconv.Itoa(i), api)
											} else {
												break
											}
										}
									} else {
										go DMTweet("Nothing."+"\n"+status.Text, api)
									}
								}
							}
						}
					case "count":
						fallthrough
					case "size":
						{
							var num int64
							var err error
							switch len(classification) {
							case 1:
								{
									num, err = redis.Int64(c.Do("DBSIZE"))
								}
							case 2:
								num, err = redis.Int64(c.Do("LLEN", strings.ToLower(classification[1])))
							}
							if err == nil {
								go DMTweet(strconv.FormatInt(num, 10)+"\n"+status.Text, api)
							} else {
								go DMTweet("err\n"+status.Text, api)
							}
						}
					}
				}
			default:
				//fmt.Println(status)
			}
		}
	}
}
