package quick

// 快速排序
func quickSort(a []int, lo, hi int) {
	if lo >= hi {
		return
	}
	p := partition(a, lo, hi)
	quickSort(a, lo, p-1)
	quickSort(a, p+1, hi)
}

func partition(a []int, lo, hi int) int {
	pivot := a[hi]
	i := lo
	for j := lo; j <= hi; j++ {
		if a[j] < pivot {
			a[j], a[i] = a[i], a[j]
			i++
		}
	}
	a[i], a[hi] = a[hi], a[i]
	return i
}

// todo understand
func quickSort2(a []int, left, right int) {
	temp := a[left]
	p := left
	i, j := left, right

	for i <= j {
		for j >= p && a[j] >= temp {
			j--
		}
		if j >= p {
			a[p] = a[j]
			p = j
		}

		for i <= p && a[i] <= temp {
			i++
		}
		if i <= p {
			a[p] = a[i]
			p = i
		}

	}


	a[p] = temp
	if p-left > 1 {
		quickSort(a, left, p-1)
	}
	if right-p > 1 {
		quickSort(a, p+1, right)
	}
}