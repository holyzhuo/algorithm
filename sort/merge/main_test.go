package merge

import (
	"testing"
	"fmt"
	"algorithm/sort"
)

var TestData = map[string][]int{
	"1": {9, 8, 1, 5, 4, 7, 6, 2, 3},
	"2": {1, 8, 7, 5, 4, 3, 2, 9, 6},
}

func TestSort(t *testing.T) {
	tr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, v := range TestData {
		res := SortBy(v)
		if !sort.IntSliceEqual(res, tr) {
			fmt.Println("test error")
		}

		fmt.Printf("%v", res)
		println()
	}
}
