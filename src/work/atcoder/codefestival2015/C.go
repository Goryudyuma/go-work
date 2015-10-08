package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	N, T := nextInt(), nextInt()
	sub := make([]int, N)
	for i := 0; i < N; i++ {
		a, b := nextInt(), nextInt()
		sub[i] = b - a
		T -= a
	}
	sort.Ints(sub)
	if T >= 0 {
		fmt.Println(0)
		return
	}
	for i := 0; i < N; i++ {
		T -= sub[i]
		if T >= 0 {
			fmt.Println(i + 1)
			return
		}
	}
	fmt.Println(-1)
}
