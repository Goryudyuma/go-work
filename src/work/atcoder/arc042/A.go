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
	N, M := nextInt(), nextInt()
	data := make([]int, M)
	for i := 0; i < M; i++ {
		data[i] = nextInt()
		data[i]--
	}
	already := make([]bool, N)
	for i := M - 1; i >= 0; i-- {
		if !already[data[i]] {
			fmt.Println(data[i] + 1)
			already[data[i]] = true
		}
	}
	for i := 0; i < N; i++ {
		if !already[i] {
			fmt.Println(i + 1)
		}
	}
}
