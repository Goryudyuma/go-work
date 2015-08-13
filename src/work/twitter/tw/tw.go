package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"math"
	"os"
	"runtime"
	"strings"
	"sync"
	"unicode/utf8"
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

//ツイート関数。無駄に並列化している。140字以上の分割ツイートにも対応。
func tweet(api *anaconda.TwitterApi, waitGroup *sync.WaitGroup, begin string, end string, str string) {
	s_len := 140 - utf8.RuneCountInString(begin) - utf8.RuneCountInString(end)
	for str != "" {
		limit := int(math.Min(float64(s_len), float64(utf8.RuneCountInString(str))))
		waitGroup.Add(1)
		go func(api *anaconda.TwitterApi, waitGroup *sync.WaitGroup, begin string, end string, str string) {
			defer waitGroup.Done()
			api.PostTweet(begin+str+end, nil)
		}(api, waitGroup, begin, end, string([]rune(str)[:limit]))
		str = string([]rune(str)[limit:])
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

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

	if utf8.RuneCountInString(begin)+utf8.RuneCountInString(end) >= 140 {
		fmt.Println("beginとendを合わせた文字列が138字未満になるようにしてください。")
		return
	}
	api := <-api_chan

	var waitGroup sync.WaitGroup
	if *c {
		fmt.Println("ようこそ連投モードへ！")
		sscan := bufio.NewScanner(os.Stdin)
		for sscan.Scan() {
			str := sscan.Text()
			tweet(api, &waitGroup, begin, end, str)
		}
		fmt.Println("ご利用ありがとうございました！")
	} else {
		tweet(api, &waitGroup, begin, end, strings.Join(flag.Args(), " "))
	}
	waitGroup.Wait()
}
