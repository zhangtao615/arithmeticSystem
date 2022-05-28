package Search

import "math/rand"

func GeneratorOrderedArray(n int) []int {
	var res []int
	for i := 0; i < n; i++ {
		res = append(res, i)
	}

	return res
}

func RandomArray(n int) []int {
	var res []int
	for i := 0; i < n; i++ {
		res = append(res, rand.Intn(100000))
	}

	return res
}
