package Search

func GeneratorOrderedArray(n int) []int {
	var res []int
	for i := 0; i < n; i++ {
		res = append(res, i)
	}

	return res
}
