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
	strs := readStrings()
	S := strs[0]
	t, _ := strconv.Atoi(strs[1])
	u, _ := strconv.Atoi(strs[2])
	if t < u {
		t, u = u, t
	}
	S = S[:t] + S[t+1:]
	if t != u {
		S = S[:u] + S[u+1:]
	}
	fmt.Println(S)
}
