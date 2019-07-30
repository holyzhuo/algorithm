package selection

// 希尔排序
func SortBy(data [9]int) [9]int {
	length := len(data)
	for gap := length / 2; gap > 0; gap = gap / 2 {
		for i := gap; i < length; i++ {
			for j := i; j >= gap; j -= gap {
				if data[j] < data[j-gap] {
					data[j], data[j-gap] = data[j-gap], data[j]
				}
			}
		}
	}

	return data
}
