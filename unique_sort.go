package main

import (
	"fmt"
)

// 排序去重
func main() {
	var array = []int{20, 40, 32, 67, 40, 20, 89, 300, 400, 15}
	var Reversal_array []int
	// 排序
	quickSort(array)
	fmt.Println(array)

	for _, v := range array {
		fmt.Println(v)
	}
	fmt.Println(Reversal_array)
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
