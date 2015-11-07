//30点解法

package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
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

// An Item is something we manage in a priority queue.
type Item struct {
	value    []int // The value of the item; arbitrary.
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
func (pq *PriorityQueue) update(item *Item, value []int, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func main() {
	W, H, Q := nextInt(), nextInt(), nextInt()
	Wboard := make([][]int64, 1e5+2)
	Hboard := make([][]int64, 1e5+2)
	Wpq := make(PriorityQueue, 0)
	Hpq := make(PriorityQueue, 0)
	for i := 0; i < 1e5+2; i++ {
		Wboard[i] = make([]int64, W)
		for j := 0; j < W; j++ {
			if i != 0 {
				Wboard[i][j] = math.MaxInt64 / 6
			} else {
				item := &Item{
					value:    []int{i, j},
					priority: 0,
				}
				heap.Push(&Wpq, item)
			}
		}
		Hboard[i] = make([]int64, H)
		for j := 0; j < H; j++ {
			if i != 0 {
				Hboard[i][j] = math.MaxInt64 / 6
			} else {
				item := &Item{
					value:    []int{i, j},
					priority: 0,
				}
				heap.Push(&Hpq, item)
			}
		}
	}
	tmax := 0
	for i := 0; i < Q; i++ {
		T, D, X := nextInt(), nextInt(), nextInt()
		if D == 0 {
			Wboard[T][X-1] = -1
			if tmax < T {
				tmax = T
			}
		} else {
			Hboard[T][X-1] = -1
			if tmax < T {
				tmax = T
			}
		}
	}
	dy := []int{0, 1, 0}
	Wans := -1
	for Wpq.Len() > 0 {
		item := heap.Pop(&Wpq).(*Item)
		nowy := item.value[0]
		if nowy == tmax {
			Wans = item.priority
			break
		}
		nowx := item.value[1]
		for k := 0; k < 3; k++ {
			nexty := nowy + dy[k]
			for l := 0; l < W; l++ {
				nextx := l
				nextcost := int(Wboard[nowy][nowx])
				nextcost += int(math.Abs(float64(nextx - nowx)))
				if nextx < W && nextx >= 0 && nexty <= tmax {
					if Wboard[nexty][nextx] > int64(nextcost) {
						Wboard[nexty][nextx] = int64(nextcost)
						itempush := &Item{
							value:    []int{nexty, nextx},
							priority: nextcost,
						}
						heap.Push(&Wpq, itempush)
					}
				}
			}
		}
	}
	if Wans == -1 {
		fmt.Println(-1)
		os.Exit(0)
	}
	Hans := -1
	for Hpq.Len() > 0 {
		item := heap.Pop(&Hpq).(*Item)
		nowy := item.value[0]
		if nowy == tmax {
			Hans = item.priority
			break
		}
		nowx := item.value[1]
		for k := 0; k < 3; k++ {
			nexty := nowy + dy[k]
			for l := 0; l < H; l++ {
				nextx := l
				nextcost := item.priority
				nextcost += int(math.Abs(float64(nextx - nowx)))
				if nextx < H && nextx >= 0 && nexty <= tmax {
					if Hboard[nexty][nextx] > int64(nextcost) {
						Hboard[nexty][nextx] = int64(nextcost)
						itempush := &Item{
							value:    []int{nexty, nextx},
							priority: nextcost,
						}
						heap.Push(&Hpq, itempush)
					}
				}
			}
		}
	}
	if Hans == -1 {
		fmt.Println(-1)
		os.Exit(0)
	}
	fmt.Println(Hans + Wans)
}
