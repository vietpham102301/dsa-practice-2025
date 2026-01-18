package algorithms

func CountingSort(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}

	// 1. Tìm min và max để xác định phạm vi
	maxVal, minVal := nums[0], nums[0]
	for _, num := range nums {
		if num > maxVal {
			maxVal = num
		}
		if num < minVal {
			minVal = num
		}
	}

	// 2. Khởi tạo mảng count với kích thước tối ưu
	k := maxVal - minVal + 1
	count := make([]int, k)

	// 3. Đếm tần suất (Sử dụng offset là minVal)
	for _, num := range nums {
		count[num-minVal]++
	}

	// 4. Cộng dồn để xác định vị trí
	for i := 1; i < k; i++ {
		count[i] += count[i-1]
	}

	// 5. Xây dựng mảng kết quả (Duyệt ngược để đảm bảo tính ổn định)
	res := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		num := nums[i]
		res[count[num-minVal]-1] = num
		count[num-minVal]--
	}

	return res
}
