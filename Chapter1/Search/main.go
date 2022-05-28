// 线性查找法 时间复杂度：O(n) 空间复杂度：O(1)
package main

import (
	"algorithmSystem/Chapter1/Search/LinerSearch"
	"fmt"
	"time"
)

func main() {
	nums := LinerSearch.GeneratorOrderedArray(100000000)
	start := time.Now()
	LinerSearch.LinerSearch(nums, 100000000-1)
	end := time.Since(start)
	fmt.Println("函数执行时间", end)
}
