package sort_algorithm

func HeapSort(nums []int) []int {
	l := len(nums)
	if l < 2 {
		return nums
	}
	// 构建最大堆
	buildMaxHeap(nums)
	// 从堆顶取出最大的元素，堆顶元素放到最后，然后重新堆化
	for i := l - 1; i >= 0; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		heapify(nums, i, 0)
	}

	return nums
}

// 构建最大堆
func buildMaxHeap(nums []int) {
	l := len(nums)
	// 从最后一个非叶子节点开始，倒序遍历
	for i := l/2 - 1; i >= 0; i-- {
		heapify(nums, l, i)
	}
}

// 堆化
func heapify(nums []int, heapSize int, index int) {
	// 最大元素位置
	largest := index
	// 左子节点
	left := 2*index + 1
	// 右子节点
	right := 2*index + 2
	// 左节点元素大于最大元素
	if left < heapSize && nums[largest] < nums[left] {
		largest = left
	}
	// 右节点元素大于最大元素
	if right < heapSize && nums[largest] < nums[right] {
		largest = right
	}
	// 递归
	if largest != index {
		// 交换最大元素和当前元素
		nums[index], nums[largest] = nums[largest], nums[index]
		// 递归处理受影响的子树
		heapify(nums, heapSize, largest)
	}
}
