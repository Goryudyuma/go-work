package main

import (
	"bufio"
	"container/list"
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
	S, G := nextInt(), nextInt()
	S--
	G--
	M := nextInt()
	cost := make([]int, N)
	order := make([]int, N)
	for i := 0; i < N; i++ {
		cost[i] = 1 << 60
		order[i] = 0
	}
	cost[S] = 0
	order[S] = 1
	route := make([][]bool, N)
	for i := 0; i < N; i++ {
		route[i] = make([]bool, N)
	}
	for i := 0; i < M; i++ {
		A, B := nextInt(), nextInt()
		A--
		B--
		route[A][B] = true
		route[B][A] = true
	}
	q := list.New()
	q.PushBack(S)
	for q.Len() > 0 {
		now, _ := q.Remove(q.Front()).(int)
		for i := 0; i < N; i++ {
			if route[now][i] {
				if cost[i] > cost[now]+1 {
					cost[i] = cost[now] + 1
					q.PushBack(i)
					order[i] = order[now]
				} else if cost[i] == cost[now]+1 {
					order[i] += order[now]
				}
				order[i] %= 1e9 + 7
			}
		}
	}
	fmt.Println(order[G])
}
