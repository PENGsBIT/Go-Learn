package main

import (
	"fmt"
)

func sortArray(nums []int) []int {
	heapSort(nums)
	return nums
}

func heapSort(nums []int) []int {
	end := len(nums) - 1
	buildMaxHeap(nums)
	for i := end; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		end--
		sink(nums, 0, end)
	}
	return nums
}
func buildMaxHeap(nums[]int)  {
	end := len(nums) - 1
	for i := end / 2; i >= 0; i-- {
		sink(nums, i, end)
	}
}
func sink(heap []int, root, end int) {
	for {
		child := root * 2 + 1
		if child > end {
			return
		}
		if child < end && heap[child] <= heap[child + 1] {
			child++
		}
		if heap[root] > heap[child] {
			return
		}
		heap[root], heap[child] = heap[child], heap[root]
		root = child
	}
}

func  quickSort(nums []int, left, right int) {
	if left > right {
		return
	}
	i, j, base := left, right, nums[left]
	for i < j {
		for nums[j] >= base && i < j {
			j--
		}
		for nums[i] <= base && i < j {
			i++
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[i], nums[left] = nums[left], nums[i]
	quickSort(nums, left, i - 1)
	quickSort(nums, i + 1, right)
}

func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2
	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, len(left) + len(right))
	l, r, i := 0, 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result[i] = left[l]
			l++
		} else {
			result[i] = right[r]
			r++
		}
		i++
	}
	copy(result[i:], left[l:])
	copy(result[i+len(left)-l:], right[r:])
	return result
}

type Person struct {
	Name string
	Age  int64
}

func main()  {
	per := Person{
		"asong",
		int64(8)
	}
	fmt.Printf("原始struct地址是：%p\n",&per)
	modifiedAge(per)
	fmt.Println(per)
}

func modifiedAge(per Person)  {
	fmt.Printf("函数里接收到struct的内存地址是：%p\n",&per)
	per.Age = 10
}




