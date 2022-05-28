// 排序算法
package main

import (
	"algorithmSystem/Chapter1/LinearSearch/Search"
	"algorithmSystem/Chapter1/Sort/SortDict"
	"fmt"
	"time"
)

func main() {
	//var nums = []int{1, 5, 4, 7, 3, 9}
	var arr = Search.RandomArray(10000)
	start := time.Now()
	SortDict.SelectionSort(arr)
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			fmt.Println("未排序成功")
			return
		}
	}
	end := time.Since(start)
	fmt.Println("排序成功, 排序花费时间为：", end)
}
