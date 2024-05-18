package sort_algorithm

func MergeSort(nums []int) []int {
	l := len(nums)
	if l < 2 {
		return nums
	}
	mid := l / 2
	left := MergeSort(nums[:mid])
	right := MergeSort(nums[mid:])

	return doMerge(left, right)
}

// 合并
func doMerge(left, right []int) []int {
	results := make([]int, 0, len(left)+len(right))
	// 记录左右数组的索引
	leftIndex, rightIndex := 0, 0
	leftLength, rightLength := len(left), len(right)
	// 左右数组都有数据时
	for leftIndex < leftLength && rightIndex < rightLength {
		// 左边数组的值小于右边数组的值，则将左边数组的值添加到结果数组中，并将左边数组的索引加1
		if left[leftIndex] < right[rightIndex] {
			results = append(results, left[leftIndex])
			leftIndex++
		} else {
			results = append(results, right[rightIndex])
			rightIndex++
		}
	}
	// 如果左边数组还有数据，则将左边数组的值添加到结果数组中
	if leftIndex < leftLength {
		results = append(results, left[leftIndex:]...)
	}
	// 如果右边数组还有数据，则将右边数组的值添加到结果数组中
	if rightIndex < rightLength {
		results = append(results, right[rightIndex:]...)
	}

	return results
}
