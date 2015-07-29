package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type point struct {
	x, y float64
}

func Add(p1 *point, p2 point) {
	p1.x += p2.x
	p1.y += p2.y
}

var sc = bufio.NewScanner(os.Stdin)

func nextInt() float64 {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return float64(i)
}

func main() {
	var N int
	sc.Split(bufio.ScanWords)
	fmt.Scan(&N)
	dataA := make([]point, N)
	dataB := make([]point, N)
	var aveA, aveB point
	for i := 0; i < N; i++ {
		dataA[i].x = nextInt()
		dataA[i].y = nextInt()
		Add(&aveA, dataA[i])
	}
	for j := 0; j < N; j++ {
		dataB[j].x = nextInt()
		dataB[j].y = nextInt()
		Add(&aveB, dataB[j])
	}
	X := float64(N)
	aveA.x /= X
	aveA.y /= X
	aveB.x /= X
	aveB.y /= X
	lenA := float64(0)
	lenB := float64(0)
	for i := 0; i < N; i++ {
		if lenA < math.Hypot(dataA[i].x-aveA.x, dataA[i].y-aveA.y) {
			lenA = math.Hypot(dataA[i].x-aveA.x, dataA[i].y-aveA.y)
		}
		if lenB < math.Hypot(dataB[i].x-aveB.x, dataB[i].y-aveB.y) {
			lenB = math.Hypot(dataB[i].x-aveB.x, dataB[i].y-aveB.y)
		}
	}
	fmt.Println(lenB / lenA)
}
