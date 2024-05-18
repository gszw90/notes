package sort_algorithm

func QuickSort(nums []int, low, high int) []int {
	l := len(nums)
	if l <= 1 {
		return nums
	}
	if low < high {
		// 获取基准元素的位置
		p := partition(nums, low, high)
		// 处理基准元素左右两侧的元素
		QuickSort(nums, low, p-1)
		QuickSort(nums, p+1, high)
	}
	return nums
}

// 分区,获取基准元素的位置
func partition(nums []int, low, high int) int {
	// 设置基准
	pivot := nums[low]
	// 小于基准元素的索引
	index := low
	for i := low + 1; i <= high; i++ {
		// 如果当前元素小于基准元素，则交换,基准元素索引后移一位
		if nums[i] <= pivot {
			index++
			nums[index], nums[i] = nums[i], nums[index]

		}
	}
	// 将基准元素交换到对应位置
	nums[index], nums[low] = nums[low], nums[index]
	return index
}
