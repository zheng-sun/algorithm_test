package main

import (
	"fmt"
)

var nums = []int{83, 34, 22, 45, 16, 90, 50, 87}

// 快速排序
func main() {
	fmt.Println("快速排序开始")
	fmt.Println("排序前:", nums)
	quickSort(nums)
	fmt.Println("排序后:", nums)
}

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

// func quickSort(nums []int) {
// 	recursionSort(nums, 0, len(nums)-1)
// }

// func recursionSort(nums []int, left int, right int) {
// 	if left < right {
// 		pivot := partition(nums, left, right)
// 		recursionSort(nums, left, pivot-1)
// 		recursionSort(nums, pivot+1, right)
// 	}
// }

// func partition(nums []int, left int, right int) int {
// 	for left < right {
// 		for left < right && nums[left] <= nums[right] {
// 			right--
// 		}
// 		if left < right {
// 			// 交换
// 			nums[left], nums[right] = nums[right], nums[left]
// 			left++
// 		}
// 		for left < right && nums[left] <= nums[right] {
// 			left++
// 		}
// 		if left < right {
// 			// 交换
// 			nums[left], nums[right] = nums[right], nums[left]
// 			right--
// 		}
// 	}
// 	return left
// }
