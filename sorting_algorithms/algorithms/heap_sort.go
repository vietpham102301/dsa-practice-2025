package algorithms

func HeapSort(nums []int) []int {
	// heapify for last node that have leaf nodes
	n := len(nums)
	for i := n/2 - 1; i >= 0; i-- {
		heapify(nums, n, i)
	}

	//sort the nums array
	for i := n - 1; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapify(nums, i, 0)
	}

	return nums
}

func heapify(nums []int, n, i int) {
	for {
		largest := i
		left, right := i*2+1, i*2+2
		if left < n && nums[left] > nums[largest] {
			largest = left
		}

		if right < n && nums[right] > nums[largest] {
			largest = right
		}

		if largest == i {
			break // heap's correct
		}

		nums[i], nums[largest] = nums[largest], nums[i]
		i = largest // continuing heapify for effected node
	}
}
