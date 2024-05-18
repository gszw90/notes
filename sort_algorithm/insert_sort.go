package sort_algorithm

func InsertSort(nums []int) []int {
	l := len(nums)
	if l < 2 {
		return nums
	}
	for i := 1; i < l; i++ {
		// 上一个数位置
		preIndex := i - 1
		//当前i位置的值
		val := nums[i]
		for preIndex >= 0 {
			// 上一个位置的数比当前位置的小，说明已经是有序的
			if nums[preIndex] <= val {
				break
			}
			// 上一个位置的数比当前位置的数大，就把上一个位置的数向后移动一位，把当前位置指针向前移动一位
			nums[preIndex+1] = nums[preIndex]
			nums[preIndex] = val
			preIndex--
		}
	}

	return nums
}
