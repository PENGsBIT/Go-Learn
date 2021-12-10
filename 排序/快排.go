package 排序

func quicksort(nums []int, l, r int) {
	if l > r {
		return
	}
	i, j, index := l, r, nums[l]
	for i < j {
		for i < j && nums[j] >= nums[index] {
			j--
		}
		for i < j && nums[i] <= nums[index] {
			i++
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[index], nums[i] = nums[i], nums[index]
	quicksort(nums, l, i-1)
	quicksort(nums, i+1, r)
}
