package 排序

func heapSort(nums []int) []int {
	if nums == nil || len(nums) == 1 {
		return nums
	}

	end := len(nums) - 1
	buildMaxHeap(nums) // 第一次排序，构建最大堆，只保证了堆顶元素是数组里最大的
	for i := end; i >= 0; i-- {
		// 这个是什么意思呢？，经过上面的一些列操作，目前array[0]是当前数组里最大的元素，需要和末尾的元素交换
		// 然后，拿出最大的元素
		nums[0], nums[i] = nums[i], nums[0]
		// 交换完后，下次遍历的时候，就应该跳过最后一个元素，也就是最大的那个值，然后开始重新构建最大堆
		// 堆的大小就减去1，然后从0的位置开始最大堆
		sink(nums, 0, i-1)
	}
	return nums
}
func buildMaxHeap(nums []int) {
	//从第一个非叶子结点从下至上，从右至左调整结构
	end := len(nums) - 1
	for i := end / 2; i >= 0; i-- {
		sink(nums, i, end)
	}
}
func sink(heap []int, root, end int) {
	for {
		child := root*2 + 1 //从root结点的左子结点开始，也就是2root+1处开始
		if child > end {
			//左子节点为空不用调整
			return
		}
		if child < end && heap[child] <= heap[child+1] {
			//如果左子结点小于右子结点，child指向右子结点
			child++
		}
		if heap[root] > heap[child] {
			return
		} else {
			//如果子节点大于父节点，将子节点值赋给父节点（不用进行交换）
			heap[root], heap[child] = heap[child], heap[root]
			root = child
		}
	}
}
