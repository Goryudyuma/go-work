package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	_, A, B := readLine(), readInts(), readInts()
	sort.Sort(sort.Reverse(sort.IntSlice(A)))
	sort.Sort(sort.Reverse(sort.IntSlice(B)))
	j := 0
	for i := range B {
		for j < len(A) && A[j] < B[i] {
			j++
		}
		j++
	}
	if j > len(A) {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}
}
