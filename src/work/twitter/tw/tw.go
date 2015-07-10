package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"os"
	"strings"
)

//認証関数
func access(fp *os.File) *anaconda.TwitterApi {
	scanner := bufio.NewScanner(fp)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
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
	return anaconda.NewTwitterApi(accesstoken, accesstoken_secret)
}

func main() {

	api_chan := make(chan *anaconda.TwitterApi)

	go func() {
		fp, err := os.Open("/etc/pass/twi.txt")
		if err != nil {
			panic(err)
		}
		defer fp.Close()

		api_chan <- access(fp)
	}()

	//引数処理
	var b_string = flag.String("b", "", "ツイートの最初に添加する文字列")
	var e_string = flag.String("e", "", "ツイートの最後に添加する文字列")
	var c = flag.Bool("c", false, "連投モードに突入")
	flag.Parse()

	begin := *b_string + " "
	end := " " + *e_string

	api := <-api_chan
	if *c {
		fmt.Println("ようこそ連投モードへ！")
		sscan := bufio.NewScanner(os.Stdin)
		for sscan.Scan() {
			str := sscan.Text()
			api.PostTweet(begin+str+end, nil)
		}
		fmt.Println("ご利用ありがとうございました！")
	} else {
		api.PostTweet(begin+strings.Join(flag.Args(), " ")+end, nil)
	}
}
