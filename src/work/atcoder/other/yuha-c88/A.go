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
	for i := 0; i < N; i++ {
		P, _, R := nextString(), nextString(), nextString()
		switch P {
		case "BEGINNING":
			fmt.Printf("%c", R[0])
		case "END":
			fmt.Printf("%c", R[len(R)-1])
		case "MIDDLE":
			fmt.Printf("%c", R[len(R)/2])
		}
	}
	fmt.Println()
}
