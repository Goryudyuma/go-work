package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var N, T, A, ans, S, G int
	ans = 0
	N = nextInt()
	T = nextInt()
	for i := 0; i < N; i++ {
		A = nextInt()
		if G < A {
			ans += G - S
			S = A
			G = A + T
		} else {
			G = A + T
		}
	}
	ans += G - S
	fmt.Println(ans)
}

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Split(bufio.ScanWords)
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}
