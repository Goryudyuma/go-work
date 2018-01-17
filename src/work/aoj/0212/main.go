package main

import (
	"container/heap"
	"fmt"
)

type Node struct {
	NowPosition          int
	DiscountTicketTCount int
	Amount               int
}

type PriorityQueueNode []*Node

func (pq PriorityQueueNode) Len() int {
	return len(pq)
}

func (pq PriorityQueueNode) Less(i, j int) bool {
	return pq[i].Amount < pq[j].Amount
}

func (pq PriorityQueueNode) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueueNode) Push(node interface{}) {
	*pq = append(*pq, node.(*Node))
}

func (pq *PriorityQueueNode) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

type Point struct {
	NextPoint int
	Fee       int
}

type Memo struct {
	NowPoint             int
	DiscountTicketTCount int
}

func main() {
	var c, n, m, s, d int
	var a, b, f int

	for {
		fmt.Scan(&c, &n, &m, &s, &d)
		if c == 0 && n == 0 && m == 0 && s == 0 && d == 0 {
			break
		}
		s--
		d--
		pointMap := make([][]Point, n)
		for i := 0; i < n; i++ {
			pointMap[i] = make([]Point, 0, n)
		}
		for i := 0; i < m; i++ {
			fmt.Scan(&a, &b, &f)
			a--
			b--
			pointMap[a] = append(pointMap[a], Point{NextPoint: b, Fee: f})
			pointMap[b] = append(pointMap[b], Point{NextPoint: a, Fee: f})
		}
		pq := make(PriorityQueueNode, 0, m)
		start := Node{NowPosition: s, DiscountTicketTCount: c, Amount: 0}
		heap.Push(&pq, &start)
		memo := make(map[Memo]int)
		for len(pq) != 0 {
			now := heap.Pop(&pq).(*Node)
			if now.NowPosition == d {
				fmt.Println(now.Amount)
				break
			}
			memo[Memo{NowPoint: now.NowPosition, DiscountTicketTCount: now.DiscountTicketTCount}] = now.Amount
			for _, v := range pointMap[now.NowPosition] {
				next := &Node{v.NextPoint, now.DiscountTicketTCount, now.Amount + v.Fee}
				if memoAmount, ok := memo[Memo{NowPoint: v.NextPoint, DiscountTicketTCount: now.DiscountTicketTCount}]; !ok || (ok && memoAmount > next.Amount) {
					memo[Memo{NowPoint: v.NextPoint, DiscountTicketTCount: now.DiscountTicketTCount}] = next.Amount
					heap.Push(&pq, next)
				}
				if now.DiscountTicketTCount > 0 {
					next = &Node{v.NextPoint, now.DiscountTicketTCount - 1, now.Amount + v.Fee/2}
					if memoAmount, ok := memo[Memo{NowPoint: v.NextPoint, DiscountTicketTCount: now.DiscountTicketTCount}]; !ok || (ok && memoAmount > next.Amount) {
						memo[Memo{NowPoint: v.NextPoint, DiscountTicketTCount: now.DiscountTicketTCount}] = next.Amount
						heap.Push(&pq, next)
					}
				}
			}
		}
	}
}
