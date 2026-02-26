package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(val any) {
	*h = append(*h, val.(int))
}

func (h *IntHeap) Pop() any {
	l := len(*h) - 1
	res := (*h)[l]

	*h = (*h)[:l]
	return res
}

func findKthLargest(nums []int, k int) int {
	h := &IntHeap{}
	heap.Init(h)

	for _, num := range nums {
		heap.Push(h, num)
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	return (*h)[0]
}

func main() {
	test := []int{1, 2, 3, 4, 5, 6, 7}

	kth := findKthLargest(test, 1)
	fmt.Println(kth)
}
