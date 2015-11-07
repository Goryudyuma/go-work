package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	S := nextString()
	fmt.Println(strconv.Itoa(strings.Count(S, "A")) + " " + strconv.Itoa(strings.Count(S, "B")) + " " + strconv.Itoa(strings.Count(S, "C")) + " " + strconv.Itoa(strings.Count(S, "D")) + " " + strconv.Itoa(strings.Count(S, "E")) + " " + strconv.Itoa(strings.Count(S, "F")))
}
