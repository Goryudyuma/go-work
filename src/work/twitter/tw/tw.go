package main

import (
	"bufio"
	"fmt"
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

	begin := ""
	end := ""
	c := 0
	n := 1
	for ; n < len(os.Args); n++ {
		if os.Args[n][0] == '-' {
			switch os.Args[n][1:] {
			case "c":
				c = 1
			case "b":
				n++
				if n < len(os.Args) {
					begin = os.Args[n] + " "
				}
			case "e":
				n++
				if n < len(os.Args) {
					end = " " + os.Args[n]
				}
			}
		} else {
			break
		}
	}
	if c == 1 {
		fmt.Println("ようこそ連投モードへ！")
		sscan := bufio.NewScanner(os.Stdin)
		for sscan.Scan() {
			str := sscan.Text()
			api.PostTweet(begin+str+end, nil)
		}
		fmt.Println("ご利用ありがとうございました！")
	} else {
		if n < len(os.Args) {
			api.PostTweet(begin+strings.Join(os.Args[n:], " ")+end, nil)
		} else {
			fmt.Println("何か入力してください。")
		}
	}
}
