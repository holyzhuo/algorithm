package main

import (
	"testing"
	"fmt"
)

var TestData = map[string][9]int{
	"1": {9, 8, 1, 5, 4, 7, 6, 2, 3},
	"2": {1, 8, 7, 5, 4, 3, 2, 9, 6},
}

func TestSort(t *testing.T) {
	tr := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, v := range TestData {
		res := SortBy(v)
		if res != tr {
			fmt.Println("test error")
		}

		fmt.Printf("%v", res)
		println()
	}
}
