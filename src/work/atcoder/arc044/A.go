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
	if N == 2 || N == 3 || N == 5 {
		fmt.Println("Prime")
		os.Exit(0)
	}
	if N == 1 || N%2 == 0 || N%5 == 0 {
		fmt.Println("Not Prime")
		os.Exit(0)
	}
	c := 0
	for N != 0 {
		c += N % 10
		N /= 10
	}
	if c%3 == 0 {
		fmt.Println("Not Prime")
		os.Exit(0)
	}
	fmt.Println("Prime")
}
