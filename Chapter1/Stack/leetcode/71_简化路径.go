// https://leetcode.cn/problems/simplify-path/  71. 简化路径

package leetcode

import (
	"strings"
)

func SimplifyPath(path string) string {
	var stack []string
	temp := map[string]int{
		"":   1,
		".":  2,
		"..": 3,
	}
	arr := strings.Split(path, "/")
	for i := 0; i < len(arr); i++ {
		if _, ok := temp[arr[i]]; !ok {
			stack = append(stack, arr[i])
		} else if len(stack) != 0 && arr[i] == ".." {
			stack = stack[0 : len(stack)-1]
		}
	}
	return "/" + strings.Join(stack, "/")
}
