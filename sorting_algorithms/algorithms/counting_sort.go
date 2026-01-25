package algorithms

func CountingSort(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}

	maxVal, minVal := nums[0], nums[0]
	for _, num := range nums {
		if num > maxVal {
			maxVal = num
		}
		if num < minVal {
			minVal = num
		}
	}

	k := maxVal - minVal + 1
	count := make([]int, k)

	for _, num := range nums {
		count[num-minVal]++
	}

	for i := 1; i < k; i++ {
		count[i] += count[i-1]
	}

	res := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		num := nums[i]
		res[count[num-minVal]-1] = num
		count[num-minVal]--
	}

	return res
}
