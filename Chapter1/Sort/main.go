// 排序算法
package main

import (
	"algorithmSystem/Chapter1/Sort/SortDict"
	"fmt"
)

func main() {
	//var nums = []int{1, 5, 4, 7, 3, 9}
	var arr = []string{"apple", "abandon", "c", "b", "d", "w", "z"}
	SortDict.SelectionSort(arr)
	fmt.Println(arr) // [abandon apple b c d w z]
}
