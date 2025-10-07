package merge_sort

// MergeSort function
func MergeSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}

	mid := len(arr) / 2

	firstArray := arr[:mid]
	secondArray := arr[mid:]

	firstArray = MergeSort(firstArray)
	secondArray = MergeSort(secondArray)

	return merge(firstArray, secondArray)
}

// Merge performing sorting in each array
func merge(arr1, arr2 []int) []int {
	// Pre‑allocate capacity to avoid repeated allocations.
	result := make([]int, 0, len(arr1)+len(arr1))

	i, j := 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] > arr2[j] {
			result = append(result, arr2[j])
			j++
		} else {
			result = append(result, arr1[i])
			i++
		}
	}

	// either arr1 or arr2 still have elements
	result = append(result, arr1[i:]...)
	result = append(result, arr2[j:]...)

	return result
}
