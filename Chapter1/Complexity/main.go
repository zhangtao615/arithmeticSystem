package main

import (
	"algorithmSystem/Chapter1/Complexity/Example"
	"fmt"
)

func main() {
	var nums [][]int = [][]int{{1, 2, 3, 4}, {5, 6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}}
	Example.SquareComplexity(nums)

	//var num = 100
	//res := Example.LogComplexity(num)
	//fmt.Println(res)

	var num = 10
	res := Example.RadicalComplexity(num)
	fmt.Println(res)
}
