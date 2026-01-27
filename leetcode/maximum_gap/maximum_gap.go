package main

import (
	"fmt"
	"math"
)

func maximumGap(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}
	minVal, maxVal := nums[0], nums[0]

	for _, num := range nums {
		if num > maxVal {
			maxVal = num
		}
		if num < minVal {
			minVal = num
		}
	}

	// Trường hợp các số đều bằng nhau
	if minVal == maxVal {
		return 0
	}

	bucketSize := (maxVal - minVal) / (n - 1)
	if bucketSize == 0 {
		bucketSize = 1
	}
	numBuckets := (maxVal-minVal)/bucketSize + 1

	minBuckets := make([]int, numBuckets)
	maxBuckets := make([]int, numBuckets)
	hasNum := make([]bool, numBuckets)

	for i := 0; i < numBuckets; i++ {
		minBuckets[i] = math.MaxInt
		maxBuckets[i] = math.MinInt
	}

	for i := 0; i < n; i++ {
		idx := (nums[i] - minVal) / bucketSize
		hasNum[idx] = true
		if nums[i] > maxBuckets[idx] {
			maxBuckets[idx] = nums[i]
		}
		if nums[i] < minBuckets[idx] {
			minBuckets[idx] = nums[i]
		}
	}

	//pigeonhole principle
	preMax := minVal
	maxGap := 0
	for i := 0; i < numBuckets; i++ {
		if !hasNum[i] {
			continue
		}
		//current Min - prevMax
		diff := minBuckets[i] - preMax
		if diff > maxGap {
			maxGap = diff
		}
		preMax = maxBuckets[i]
	}

	return maxGap
}

func main() {
	nums := []int{3, 6, 9, 1}
	fmt.Printf("Maximum Gap của %v là: %d\n", nums, maximumGap(nums))
}
