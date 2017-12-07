package main

import (
	"fmt"
)

// 排序去重
func main() {
	var array = []int{20, 40, 32, 67, 40, 20, 89, 300, 400, 15}
	fmt.Println("原数组:", array)
	// 排序
	quickSort(array)
	fmt.Println("排序后:", array)

	// 去重
	var array_new []int
	for i := 0; i < len(array); i++ {
		if i+1 < len(array) && array[i] != array[i+1] {
			array_new = append(array_new, array[i])
		} else if i+1 == len(array) {
			array_new = append(array_new, array[i])
		}
	}
	fmt.Println("排序去重:", array_new)
}

// 排序
func quickSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	head := 0
	mid, tail := nums[0], len(nums)-1
	for i := 1; i <= tail; {
		if nums[i] > mid {
			nums[i], nums[tail] = nums[tail], nums[i]
			tail--
		} else {
			nums[i], nums[head] = nums[head], nums[i]
			head++
			i++
		}
	}
	nums[head] = mid
	quickSort(nums[:head])
	quickSort(nums[head+1:])
}
