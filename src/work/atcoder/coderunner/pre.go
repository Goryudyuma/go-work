package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func strreplace(input string, n int, char string) (ret string) {
	ret = input[:n] + char + input[n+1:]
	return
}

func point(S string) {
	response, _ := http.Get("http://sample.coderunner.jp/q?token=hntVKIWFwqAvchh2SY4NsKPyRSs4rAuM&str=" + S)
	body, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	fmt.Println(string(body))
}
func main() {
	N := 4
	var S, end string
	for i := 0; i < N; i++ {
		S = S + "A"
		end = end + "D"
	}
	now := N - 1
	fmt.Println(S)
	point(S)
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
		fmt.Println(S)

		point(S)
		time.Sleep(1000000000)
	}
}
