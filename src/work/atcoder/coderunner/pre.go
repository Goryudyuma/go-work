package main

import (
	"github.com/garyburd/redigo/redis"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

var maxpoint int
var memo string

func strreplace(input string, n int, char string) (ret string) {
	ret = input[:n] + char + input[n+1:]
	return
}

func point(S string, c redis.Conn) {
	response, _ := http.Get("http://sample.coderunner.jp/q?token=hntVKIWFwqAvchh2SY4NsKPyRSs4rAuM&str=" + S)
	body, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	//fmt.Println(string(body))
	x, _ := strconv.Atoi(string(body))
	c.Do("SET", S, x)
	if x > maxpoint {
		maxpoint = x
		memo = S
		ioutil.WriteFile("result.txt", []byte(memo+"\n"+strconv.Itoa(maxpoint)+"\n"), os.ModePerm)
	}
}

func main() {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		panic(err)
	}
	defer c.Close()
	for N := 1; N < 8; N++ {
		maxpoint = 0
		var S, end string
		for i := 0; i < N; i++ {
			S = S + "A"
			end = end + "D"
		}
		now := N - 1
		//fmt.Println(S)

		c.Do("SELECT", "1")
		point(S, c)
		time.Sleep(1000000000)
		for S != end {
			switch S[now] {
			case 'A':
				{
					S = S[:now] + "B" + S[now+1:]
					now = N - 1
				}
			case 'B':
				{
					S = S[:now] + "C" + S[now+1:]
					now = N - 1
				}
			case 'C':
				{
					S = S[:now] + "D" + S[now+1:]
					now = N - 1
				}
			case 'D':
				{
					S = S[:now] + "A" + S[now+1:]
					now--
					continue
				}
			}
			//fmt.Println(S)

			point(S, c)
			time.Sleep(1000000000)
		}
	}
}
