package main

// http://golang.org/pkg/container/heap/ を参照に、priority queueを作る

import (
	"bufio"
	"container/heap"
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

type point struct {
	x, y int
}

// An Item is something we manage in a priority queue.
type Item struct {
	value    point // The value of the item; arbitrary.
	priority int   // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value point, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func main() {
	H, W, T := nextInt(), nextInt(), nextInt()
	S, calcdata := make([][]byte, H), make([][]int, H)
	var start, goal point
	for i := 0; i < H; i++ {
		calcdata[i] = make([]int, W)
		S[i] = make([]byte, W)
		S[i] = []byte(nextString())
		for j := 0; j < W; j++ {
			if S[i][j] == 'S' {
				start = point{x: i, y: j}
			}
			if S[i][j] == 'G' {
				goal = point{x: i, y: j}
			}
		}
	}
	X_MIN, X_MAX := 0, 1<<60
	dx := [4]int{0, 1, 0, -1}
	dy := [4]int{1, 0, -1, 0}
	for X_MAX-X_MIN > 1 {
		for i := 0; i < H; i++ {
			for j := 0; j < W; j++ {
				calcdata[i][j] = 1 << 40
			}
		}
		calcdata[start.x][start.y] = 0
		X_MID := int((X_MIN + X_MAX) / 2)
		pq := make(PriorityQueue, 0)
		heap.Init(&pq)
		start_pq := &Item{
			value:    start,
			priority: 0,
		}
		heap.Push(&pq, start_pq)
		for pq.Len() > 0 {
			item := heap.Pop(&pq).(*Item)
			for i := 0; i < 4; i++ {
				now := point{x: item.value.x + dx[i], y: item.value.y + dy[i]}
				if 0 > now.x || H <= now.x || 0 > now.y || W <= now.y {
					continue
				}
				switch S[now.x][now.y] {
				case '.':
					{
						if calcdata[now.x][now.y] > calcdata[item.value.x][item.value.y]+1 {
							calcdata[now.x][now.y] = calcdata[item.value.x][item.value.y] + 1
							updata_item := &Item{
								value:    now,
								priority: calcdata[now.x][now.y],
							}
							heap.Push(&pq, updata_item)
							pq.update(updata_item, updata_item.value, len(pq))
						}

					}
				case 'S':
					{
						calcdata[now.x][now.y] = 0
					}
				case 'G':
					{
						if calcdata[now.x][now.y] > calcdata[item.value.x][item.value.y]+1 {
							calcdata[now.x][now.y] = calcdata[item.value.x][item.value.y] + 1
						}
					}
				case '#':
					{
						if calcdata[now.x][now.y] > calcdata[item.value.x][item.value.y]+X_MID {
							calcdata[now.x][now.y] = calcdata[item.value.x][item.value.y] + X_MID
							updata_item := &Item{
								value:    now,
								priority: calcdata[now.x][now.y],
							}
							heap.Push(&pq, updata_item)
							pq.update(updata_item, updata_item.value, len(pq))
						}
					}
				}
			}
		}
		if calcdata[goal.x][goal.y] <= T {
			X_MIN = X_MID
		} else {
			X_MAX = X_MID
		}
	}
	fmt.Println(X_MIN)
}
