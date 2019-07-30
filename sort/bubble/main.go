package main

// 冒泡排序
func SortBy(data [9]int) [9]int {
	length := len(data)
	if length == 0 {
		return data
	}

	for j := 0; j < length-1; j++ {
		for i := 0; i < length-j-1; i++ {
			if data[i] > data[i+1] {
				data[i], data[i+1] = data[i+1], data[i]
			}
		}
	}
	return data
}
