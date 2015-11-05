package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"sort"
	"strconv"
)

func get(S string, c redis.Conn) int {
	num, err := redis.String(c.Do("GET", S))
	if err != nil {
		panic(err)
	}
	ret, err := strconv.Atoi(num)
	if err != nil {
		panic(err)
	}
	return ret
}

func main() {
	N := 8

	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		panic(err)
	}
	defer c.Close()

	var S, end string
	for i := 0; i < N; i++ {
		S = S + "A"
		end = end + "D"
	}
	now := N - 1

	c.Do("SELECT", "1")
	a := List{}
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

		//SI[S] = get(S, c)

		a = append(a, Entry{S, get(S, c)})
	}

	/*	for k, v := range SI {
			e := Entry{k, v}
			a = append(a, e)
		}
	*/

	fmt.Println("!")

	sort.Sort(a)
	fmt.Println(a)
}

type Entry struct {
	name  string
	value int
}
type List []Entry

func (l List) Len() int {
	return len(l)
}

func (l List) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l List) Less(i, j int) bool {
	if l[i].value == l[j].value {
		return (l[i].name < l[j].name)
	} else {
		return (l[i].value < l[j].value)
	}
}
