package main

// 插入排序
func SortBy(data [9]int) [9]int {
	length := len(data)
	for i := 0; i < length-1; i++ {
		for j := i + 1; j > 0; j-- {
			if data[j] < data[j-1] {
				data[j], data[j-1] = data[j-1], data[j]
			}
		}
	}
	return data
}
