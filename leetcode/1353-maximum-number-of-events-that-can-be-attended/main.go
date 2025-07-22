package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func maxEvents(events [][]int) int {
	maxDay := 0
	for _, event := range events {
		maxDay = max(maxDay, event[1])
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})
	// fmt.Println(maxDay)

	minHeap := &MinHeap{}
	heap.Init(minHeap)
	ans := 0
	n := len(events)
	for i, j := 1, 0; i <= maxDay; i++ {
		for j < n && events[j][0] <= i {
			heap.Push(minHeap, events[j][1])
			j++
		}
		// remove unnecessary events
		for minHeap.Len() > 0 && (*minHeap)[0] < i {
			heap.Pop(minHeap)
		}
		if minHeap.Len() > 0 {
			heap.Pop(minHeap)
			ans++
		}
	}
	return ans
}

func main() {
	events := [][]int{{1, 4}, {4, 4}, {2, 2}, {3, 4}, {1, 1}}

	result := maxEvents(events)

	fmt.Printf("Input: %v, result: %d\n", events, result)
}
