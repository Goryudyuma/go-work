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
	data := make([]int, N)
	a, b := nextInt(), nextInt()
	a--
	b--
	data[a]++
	data[b]++
	K := nextInt()
	for i := 0; i < K; i++ {
		p := nextInt()
		p--
		data[p]++
	}
	f := true
	for i := 0; i < N; i++ {
		if data[i] > 1 {
			f = false
		}
	}
	if f {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
