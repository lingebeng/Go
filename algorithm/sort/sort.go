package sort

func QuickSort(a *[]int, l, r int) {
	arr := *a
	if l >= r {
		return
	}
	x := arr[(l+r)>>1]
	i, j := l-1, r+1
	for i < j {
		i++
		for arr[i] < x {
			i++
		}
		j--
		for arr[j] > x {
			j--
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	QuickSort(a, l, j)
	QuickSort(a, j+1, r)
}
