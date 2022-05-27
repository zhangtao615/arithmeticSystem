// 线性查找法 时间复杂度：O(n) 空间复杂度：O(1)
package main

import (
	"algorithmSystem/Chapter1/LinearSearch/Search"
	"fmt"
)

func main() {
	nums := []float32{1.1, 2.2, 3.3, 4.4, 5.5, 6.6}

	res := Search.Search(nums, 5.4)

	fmt.Printf("target's index is %d", res)
}
