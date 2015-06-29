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

func nextFloat() float64 {
	f, e := strconv.ParseFloat(nextString(), 64)
	if e != nil {
		panic(e)
	}
	return f
}

func nextString() string {
	var sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	return sc.Text()
}

func main() {
	var G, D, N int
	fmt.Scan(&N)
	f := -1
	for i := 0; i < N; i++ {
		fmt.Scan(&G)
		fmt.Scan(&D)
		if (G-30000*D)*6 >= 3000000 && f == -1 {
			f = i + 1
		}
	}
	if f == -1 {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
		for i := 0; i < 6; i++ {
			fmt.Println(f)
		}
	}
}
