package main

import "fmt"

func main() {
	var nums = []int{1, 2, 3, 4, 5, 6}
	temp := nums[len(nums)-1 : len(nums)]
	fmt.Println(temp)
	//var emptyNums = make([]int)
	//for i, v := range nums {
	//	fmt.Printf("nums的index是%d, nums[%d]=%d", i, i, v)
	//	fmt.Println()
	//}
}
