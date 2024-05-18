package sort_algorithm

func ShellSort(nums []int) []int {
	l := len(nums)
	if l < 2 {
		return nums
	}
	// 间隔长度
	gap := l
	for gap > 1 {
		// 计算生成新的间隔
		gap = gap/3 + 1
		for i := gap; i < l; i++ {
			// 记录当前值
			val := nums[i]
			// 当前值的位置
			j := i
			// 对i前面的元素根据gap来进行插入排序，
			for j >= gap && nums[j-gap] > val {
				nums[j], nums[j-gap] = nums[j-gap], nums[j]
				j -= gap
			}

		}
	}

	return nums
}
