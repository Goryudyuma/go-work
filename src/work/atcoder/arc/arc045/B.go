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

type pair struct {
	x, y int
}

func main() {
	N, M := nextInt(), nextInt()
	data := make([]pair, M)
	N += 2
	place := make([]int, N)
	for i := 0; i < M; i++ {
		data[i] = pair{nextInt(), nextInt()}
		place[data[i].x]++
		place[data[i].y+1]--
	}
	for i := 1; i < N; i++ {
		place[i] += place[i-1]
	}
	for i := 1; i < N; i++ {
		if place[i] == 1 {
			place[i] += place[i-1]
		} else {
			place[i] = place[i-1]
		}
	}
	ans := make([]int, 0)
	for i := 0; i < M; i++ {
		if place[data[i].x-1] == place[data[i].y] {
			ans = append(ans, i)
		}
	}
	fmt.Println(len(ans))
	for i := range ans {
		fmt.Println(ans[i] + 1)
	}
}
