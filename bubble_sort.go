package main

import (
	"fmt"
)

var t = [5]int{3, 68, 46, 1, 22}
var a int

// 冒泡排序
func main() {
	fmt.Println("冒泡排序开始")
	for j := 0; j < len(t)-1; j++ {
		for i := 0; i < len(t)-1; i++ {
			if t[i] > t[i+1] {
				// 交换
				t[i], t[i+1] = t[i+1], t[i]
			}
		}
	}
	fmt.Println(t)

}
