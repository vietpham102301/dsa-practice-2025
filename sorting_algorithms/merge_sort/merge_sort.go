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
	var result []int

	for len(arr1) > 0 && len(arr2) > 0 {
		if arr1[0] > arr2[0] {
			result = append(result, arr2[0])
			arr2 = arr2[1:]
		} else {
			result = append(result, arr1[0])
			arr1 = arr1[1:]
		}
	}

	// either arr1 or arr2 still have elements
	result = append(result, arr1...)
	result = append(result, arr2...)

	return result
}
