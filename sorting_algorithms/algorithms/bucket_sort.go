package algorithms

import "sort"

func BucketSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	bucketSize := 10

	minVal, maxVal := arr[0], arr[0]
	for _, v := range arr {
		if v > maxVal {
			maxVal = v
		}
		if v < minVal {
			minVal = v
		}
	}

	bucketNum := (maxVal-minVal)/bucketSize + 1

	buckets := make([][]int, bucketNum)

	//scatter
	for _, v := range arr {
		index := (v - minVal) / bucketSize
		buckets[index] = append(buckets[index], v)
	}

	res := make([]int, 0, len(arr))
	//sort bucket & gather
	for _, bucket := range buckets {
		sort.Ints(bucket)
		res = append(res, bucket...)
	}

	return res
}
