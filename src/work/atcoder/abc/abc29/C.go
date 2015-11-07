package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	i, e := strconv.Atoi(nextString())
	if e != nil {
		panic(e)
	}
	return i
}

func nextInt64() int64 {
	i, e := strconv.ParseInt(nextString(), 10, 64)
	if e != nil {
		panic(e)
	}
	return i

}

func nextFloat() float64 {
	f, e := strconv.ParseFloat(nextString(), 64)
	if e != nil {
		panic(e)
	}
	return f
}

func nextString() string {
	sc.Split(bufio.ScanWords)
	sc.Scan()
	return sc.Text()
}

func main() {
	N := nextInt()
	var S, end string
	for i := 0; i < N; i++ {
		S = S + "a"
		end = end + "c"
	}
	now := N - 1
	fmt.Println(S)
	for S != end {
		switch S[now] {
		case 'a':
			{
				S = S[:now] + "b" + S[now+1:]
				now = N - 1
			}
		case 'b':
			{
				S = S[:now] + "c" + S[now+1:]
				now = N - 1
			}
		case 'c':
			{
				S = S[:now] + "a" + S[now+1:]
				now--
				continue
			}
		}
		fmt.Println(S)
	}
}
