package SortDict

type ParamType interface {
	int | int32 | float32 | float64 | int64 | int8 | string
}

func SelectionSort[T ParamType](nums []T) {
	length := len(nums)

	for i := 0; i < length; i++ {
		// 选择 nums[i...n) 中的最小值索引
		minIndex := i

		for j := i; j < length; j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		swap(nums, i, minIndex)
	}
}

func swap[T comparable](nums []T, start, end int) {
	temp := nums[start]
	nums[start] = nums[end]
	nums[end] = temp
}
