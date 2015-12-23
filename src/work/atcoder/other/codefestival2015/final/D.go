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
	if N == 2 {
		fmt.Println(1)
		return
	}
	A := make([]int, 1e5+5)
	p, q := make([]int, N), make([]int, N)
	for i := 0; i < N; i++ {
		S, T := nextInt(), nextInt()
		A[S]++
		A[T]--
		p[i], q[i] = S, T
	}
	maxi, maxmini, maxmaxi := -1, -1, -1
	for i := 1; i < 1e5+5; i++ {
		A[i] += A[i-1]
	}
	for i := 0; i < 1e5+5; i++ {
		if maxi < A[i] {
			maxmini, maxi = i, A[i]
		}
		if maxi == A[i] {
			maxmaxi = i
		}
	}
	for i := 0; i < N; i++ {
		if p[i] <= maxmini && maxmaxi <= q[i]+1 {
			maxi--
			break
		}
	}
	fmt.Println(maxi)
}
