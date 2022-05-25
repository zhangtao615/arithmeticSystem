// 线性查找法 时间复杂度：O(n) 空间复杂度：O(1)
package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}

	res := Search(nums, 10)

	fmt.Printf("target's index is %d", res)
}

func Search(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		}
	}

	return -1
}
