package SortDict

func InsertionSort(arr []int) {
	length := len(arr)

	for i := 0; i < length; i++ {
		// 将arr[i] 插入到合适的位置
		for j := i; j-1 >= 0; j-- {
			if arr[j] < arr[j-1] {
				Swap(arr, j, j-1)
			}
		}
	}
}

func InsertionSortOptimize(arr []int) {
	length := len(arr)

	for i := 0; i < length; i++ {
		temp := arr[i]
		var j int
		for j = i; j-1 >= 0 && temp < arr[j-1]; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = temp
	}
}
