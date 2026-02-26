package main

import "fmt"

func subarraySum(nums []int, k int) int {
	prefixSum := make(map[int]int, len(nums)+1)
	prefixSum[0] = 1

	currentSum := 0
	count := 0
	for _, num := range nums {
		currentSum += num
		count += prefixSum[currentSum-k]

		prefixSum[currentSum] += 1
	}

	return count
}

func main() {
	test := []int{1, 2, 3}
	res := subarraySum(test, 3)

	fmt.Println(res)
}
