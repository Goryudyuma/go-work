package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var rdr = bufio.NewReaderSize(os.Stdin, 1000000)

func readLine() string {
	buf := make([]byte, 0, 1000000)
	for {
		l, p, _ := rdr.ReadLine()
		buf = append(buf, l...)
		if !p {
			break
		}
	}
	return string(buf)
}

func readStrings() []string {
	return strings.Split(readLine(), " ")
}

func readInts() (ret []int) {
	strings := readStrings()
	for i := range strings {
		num, _ := strconv.Atoi(strings[i])
		ret = append(ret, num)
	}
	return
}

func main() {
	NM := readInts()
	N, M := NM[0], NM[1]
	XY := readInts()
	X, Y := XY[0], XY[1]
	a := readInts()
	b := readInts()
	i, j := 0, 0
	nowtime := 0
	aorb := 0
	count := 0
	for {
		for i < N && a[i] < nowtime {
			i++
		}
		for j < M && b[j] < nowtime {
			j++
		}
		if (i == N && aorb == 0) || (j == M && aorb == 1) {
			break
		}
		if aorb == 0 {
			nowtime = a[i] + X
		} else {
			nowtime = b[j] + Y
			count++
		}
		aorb = 1 - aorb

	}

	//fmt.Println(readLine())
	//fmt.Println(readInts())
	fmt.Println(count)
}
