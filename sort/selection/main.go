package selection

func SortBy(data [9]int) [9]int {
	length := len(data)
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if data[i] > data[j] {
				data[i], data[j] = data[j], data[i]
			}
		}
	}
	return data
}
