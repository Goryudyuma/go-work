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
	N, M, S := nextInt(), nextInt(), nextInt()
	data := make([]int, 10001)
	for i := 0; i < N; i++ {
		T, K := nextInt(), nextInt()
		data[T] = K
	}
	count := 0
	for i := 1; i < 10001; i++ {
		data[i] += data[i-1]
		if i >= S {
			break
		}
		if data[i] >= M {
			count++
		}
	}
	fmt.Println(count)
}
