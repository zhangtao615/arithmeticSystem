package Search

type T interface {
}

func Search[T comparable](nums []T, target T) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		}
	}

	return -1
}
