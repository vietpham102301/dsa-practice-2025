package algorithms

func medianOfThree(arr []int, low, high int) int {
	mid := (low + high) / 2

	if arr[low] > arr[mid] {
		arr[low], arr[mid] = arr[mid], arr[low]
	}

	if arr[low] > arr[high] {
		arr[low], arr[high] = arr[high], arr[low]
	}

	if arr[mid] > arr[high] {
		arr[mid], arr[high] = arr[high], arr[mid]
	}

	// đưa median về cuối mảng
	arr[mid], arr[high] = arr[high], arr[mid]
	return partition(arr, low, high)
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]

	return i + 1
}

func QuickSort(arr []int, low, high int) {
	if low < high {
		pi := medianOfThree(arr, low, high)
		QuickSort(arr, low, pi-1)
		QuickSort(arr, pi+1, high)
	}
}
