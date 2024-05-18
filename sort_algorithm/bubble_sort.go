package sort_algorithm

func BubbleSort(nums []int) []int {
	l := len(nums)
	if l < 2 {
		return nums
	}
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}

	return nums
}
