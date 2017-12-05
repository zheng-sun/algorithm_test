package main

import (
	"fmt"
)

var buckers [10]int
var t = [5]int{5, 2, 5, 8, 7}
var i, j int

// 桶排序
func main() {
	fmt.Print("桶排序开始")

}

func bucketSort() {
	for i := 0; i < 10; i++ {
		buckers[i] = 0
	}
	fmt.Println(buckers)

	for _, v := range t {
		buckers[v]++
	}
	fmt.Println(buckers)

	for i := 0; i < 10; i++ {
		for j := 1; j <= buckers[i]; j++ {
			fmt.Println(i)
		}
	}
}
