package main

import (
	"bufio"
	"fmt"
	"os"
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

func main() {
	STU := readStrings()
	if len(STU) == 3 && len(STU[0]) == 5 && len(STU[1]) == 7 && len(STU[2]) == 5 {
		fmt.Println("valid")
	} else {
		fmt.Println("invalid")
	}
}
