package Example

import (
	"fmt"
	"strconv"
)

// 时间复杂度为O(n^2)

func SquareComplexity[T comparable](nums [][]T) {
	length := len(nums)
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			fmt.Println(nums[i][j])
		}
	}
}

// 时间复杂度为O(logN) w二进制转十进制

func LogComplexity(num int) string {
	s := ""
	if num == 0 {
		return "0"
	}
	for ; num > 0; num /= 2 {
		lsb := num % 2
		s = strconv.Itoa(lsb) + s
	}
	return s
}

// 时间复杂度为O(&radic;n) 求数字n的所有约数

func RadicalComplexity(num int) []int {
	var res []int
	if num == 0 {
		return []int{}
	}

	for i := 1; i*i <= num; i++ {
		if num%i == 0 {
			res = append(res, i, num/i)
		}
	}
	return res
}
