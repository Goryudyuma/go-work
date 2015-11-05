package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func nextInt() int {
	i, e := strconv.Atoi(nextString())
	if e != nil {
		panic(e)
	}
	return i
}

func nextString() string {
	var sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	return sc.Text()
}
func main() {
	var H, W, N, K int

	fmt.Scan(&H)
	fmt.Scan(&W)
	fmt.Scan(&N)
	fmt.Scan(&K)

	if H*W%N == K%N {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
