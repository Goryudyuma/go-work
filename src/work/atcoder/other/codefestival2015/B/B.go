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
	A := make([]int, M+1)
	for i := 0; i < N; i++ {
		x := nextInt()
		A[x]++
	}
	for i := 0; i < M+1; i++ {
		if A[i] > N/2 {
			fmt.Println(i)
			os.Exit(0)
		}
	}
	fmt.Println("?")
}
