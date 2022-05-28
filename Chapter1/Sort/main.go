// 排序算法
package main

import (
	"algorithmSystem/Chapter1/Search/LinerSearch"
	"algorithmSystem/Chapter1/Sort/SortDict"
	"fmt"
	"time"
)

func main() {
	//var nums = []int{1, 5, 4, 7, 3, 9}
	var arr = LinerSearch.RandomArray(100000)
	start := time.Now()
	//SortDict.SelectionSortReverse(arr)
	SortDict.InsertionSortOptimize(arr)
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			fmt.Println("未排序成功")
			return
		}
	}
	end := time.Since(start)
	fmt.Println("排序成功, 排序花费时间为：", end)
}
