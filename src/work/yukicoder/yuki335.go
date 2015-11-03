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
	readLine()
	S := readLine()
	for i := 0; i < len(S); i++ {
		for j := 1; j <= (len(S)-i)/2; j++ {
			if S[i:i+j] == S[i+j:i+j+i+j-i] {
				fmt.Println("YES")
				os.Exit(0)
			}
		}
	}
	fmt.Println("NO")
	os.Exit(0)
}
