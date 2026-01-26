package algorithms

import "math"

// RadixSort sorting
func RadixSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	var positive []int
	var negatives []int
	for _, v := range arr {
		if v >= 0 {
			positive = append(positive, v)
		} else {
			negatives = append(negatives, v)
		}
	}

	// Chỉ sort nếu mảng có phần tử
	if len(positive) > 0 {
		maxP := getMaxABS(positive)
		for exp := 1; maxP/exp > 0; exp *= 10 {
			countingSort(positive, exp)
		}
	}

	if len(negatives) > 0 {
		maxN := getMaxABS(negatives)
		for exp := 1; maxN/exp > 0; exp *= 10 {
			countingSort(negatives, exp)
		}
		// Đảo ngược mảng âm: |-100| > |-1| nhưng -100 < -1
		for i, j := 0, len(negatives)-1; i < j; i, j = i+1, j-1 {
			negatives[i], negatives[j] = negatives[j], negatives[i]
		}
	}

	return append(negatives, positive...)
}

// Counting sort
func countingSort(arr []int, exp int) {
	count := make([]int, 10)
	output := make([]int, len(arr))

	//count frequency
	for _, v := range arr {
		index := (int(math.Abs(float64(v))) / exp) % 10
		count[index]++
	}

	//prefix sum
	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}

	//sort shit
	for i := len(arr) - 1; i >= 0; i-- {
		index := count[(int(math.Abs(float64(arr[i])))/exp)%10] - 1
		output[index] = arr[i]
		count[int(math.Abs(float64(arr[i])))]--
	}

	//copy shit
	for i := 0; i < len(arr); i++ {
		arr[i] = output[i]
	}
}

// Get max abs
func getMaxABS(arr []int) int {
	maxVal := int(math.Abs(float64(arr[0])))
	for _, v := range arr {
		if int(math.Abs(float64(v))) > maxVal {
			maxVal = int(math.Abs(float64(v)))
		}
	}

	return maxVal
}
