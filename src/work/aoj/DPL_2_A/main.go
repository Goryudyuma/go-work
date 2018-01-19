package main

import (
	"container/heap"
	"fmt"
)

type Scanner struct{}

func (sc Scanner) Next() (s string) {
	fmt.Scan(&s)
	return
}

func (sc Scanner) NextInt() (i int) {
	fmt.Scan(&i)
	return
}

type MinCostStatus struct {
	Visited   int
	PrevPoint int
}

type Status struct {
	MinCostStatus
	Cost int
}

type PriorityQueueStatus []*Status

func (pq PriorityQueueStatus) Len() int {
	return len(pq)
}

func (pq PriorityQueueStatus) Less(i, j int) bool {
	return pq[i].Cost < pq[j].Cost
}

func (pq PriorityQueueStatus) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueueStatus) Push(data interface{}) {
	*pq = append(*pq, data.(*Status))
}

func (pq *PriorityQueueStatus) Pop() interface{} {
	l := len(*pq)
	item := (*pq)[l-1]
	*pq = (*pq)[0 : l-1]
	return item
}

func main() {
	var sc Scanner
	v, e := sc.NextInt(), sc.NextInt()
	m := make([]map[int]int, v)
	for i := 0; i < v; i++ {
		m[i] = make(map[int]int)
	}
	for i := 0; i < e; i++ {
		s, t, d := sc.NextInt(), sc.NextInt(), sc.NextInt()
		m[s][t] = d
	}
	pq := make(PriorityQueueStatus, 0)
	start := &Status{
		MinCostStatus: MinCostStatus{
			Visited:   0,
			PrevPoint: 0,
		},
		Cost: 0,
	}
	memo := make(map[MinCostStatus]int)
	heap.Push(&pq, start)
	for pq.Len() > 0 {
		now := heap.Pop(&pq).(*Status)
		if now.Visited == (1<<uint(v))-1 && now.PrevPoint == 0 {
			fmt.Println(now.Cost)
			return
		}
		for k, v := range m[now.PrevPoint] {
			if now.Visited&(1<<uint(k)) == 0 {
				minCostStatus := MinCostStatus{
					Visited:   now.Visited | (1 << uint(k)),
					PrevPoint: k,
				}
				cost := now.Cost + v
				if v, ok := memo[minCostStatus]; !ok || (ok && cost < v) {
					next := &Status{
						MinCostStatus: minCostStatus,
						Cost:          cost,
					}
					memo[minCostStatus] = cost
					heap.Push(&pq, next)
				}
			}
		}
	}
	fmt.Println(-1)
}
