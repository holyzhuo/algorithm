package main

func SortBy(data [9]int) [9]int {
    for j := 0; j < len(data); j++ {
        for i := 0; i < len(data) - j - 1; i++  {
            if data[i] > data[i+1] {
                data[i], data[i+1] = data[i+1], data[i]
            }
        }
    }
    return data
}

