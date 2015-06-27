package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	N := nextInt()
	X := 0
	for N != 0 {
		N--
		a := nextInt()
		if a%2 == 0 {
			X--
		} else {
			X++
		}
	}
	fmt.Printf("%v\n", math.Abs(float64(X)))
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
