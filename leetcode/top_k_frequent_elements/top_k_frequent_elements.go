package main

import "container/heap"

func topKFrequent(nums []int, k int) []int {
	frequencyMap := make(map[int]int, len(nums))

	for _, num := range nums {
		frequencyMap[num]++
	}

	buckets := make([][]int, len(nums)+1)
	for key, val := range frequencyMap {
		buckets[val] = append(buckets[val], key)
	}

	res := make([]int, 0, k)
	for i := len(buckets) - 1; i >= 0 && len(res) < k; i-- {
		res = append(res, buckets[i]...)
	}

	return res[:k]
}

type Item struct {
	key   int
	count int
}

type MinHeap []Item

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].count < h[j].count
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(val any) {
	(*h) = append((*h), val.(Item))
}

func (h *MinHeap) Pop() any {
	old := (*h)
	l := len(old)
	item := old[l-1]
	(*h) = (*h)[:l-1]
	return item
}

func topKFrequentHeap(nums []int, k int) []int {
	frequencyMap := make(map[int]int, len(nums))

	for _, num := range nums {
		frequencyMap[num]++
	}

	minHeap := &MinHeap{}

	for key, val := range frequencyMap {
		heap.Push(minHeap, Item{
			key:   key,
			count: val,
		})
		if minHeap.Len() > k {
			heap.Pop(minHeap)
		}
	}

	res := make([]int, k)
	for i := k - 1; i >= 0; i-- {
		res[i] = heap.Pop(minHeap).(Item).key
	}

	return res
}
