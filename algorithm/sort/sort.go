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

func MergeSort(a *[]int, l, r int) {
	arr := *a
	if l >= r {
		return
	}
	mid := (l + r) >> 1
	MergeSort(a, l, mid)
	MergeSort(a, mid+1, r)
	i, j, k := l, mid+1, 0
	tmp := make([]int, r-l+1)
	for i <= mid && j <= r {
		if arr[i] < arr[j] {
			tmp[k] = arr[i]
			i++
		} else {
			tmp[k] = arr[j]
			j++
		}
		k++
	}
	for i <= mid {
		tmp[k] = arr[i]
		k++
		i++
	}
	for j <= r {
		tmp[k] = arr[j]
		k++
		j++
	}
	i, k = l, 0
	for i <= r {
		arr[i] = tmp[k]
		i++
		k++
	}
}
