package sort_algorithm

func SelectSort(nums []int) []int {
	l := len(nums)
	if l < 2 {
		return nums
	}
	left, right := 0, l-1
	for left <= right {
		// 记录区间内最小与最大数的位置
		minIndex := left
		maxIndex := right
		for i := left; i <= right; i++ {
			// 最小数位置
			if nums[i] < nums[minIndex] {
				minIndex = i
			}
			// 最大数位置
			if nums[i] > nums[maxIndex] {
				maxIndex = i
			}
		}
		// 如果最大位置在left处，当把最小处的元素交换到left处后，
		// left处原来最大的元素就被交换到了minxIndex处
		if maxIndex == left {
			maxIndex = minIndex
		}

		// 将最小元素移动到左边
		if left != minIndex {
			nums[left], nums[minIndex] = nums[minIndex], nums[left]
		}
		// 将最大元素移动到右边
		if right != maxIndex {
			nums[right], nums[maxIndex] = nums[maxIndex], nums[right]
		}
		// 缩小范围
		left++
		right--
	}

	return nums
}
