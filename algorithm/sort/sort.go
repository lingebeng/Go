package sort

// QuickSort 快速排序，不稳定
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

// MergeSort 归并排序，稳定
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

// Heapify 保持大顶堆的性质，最大的值位于根节点
func Heapify(arr []int, n int, i int) {
	largest := i
	left, right := 2*i+1, 2*i+2 // 下 标 从 0 开 始 ！
	if left < n && arr[left] > arr[largest] {
		largest = left
	}
	if right < n && arr[right] > arr[largest] {
		largest = right
	}
	// 如果当前结点与左右儿子结点比不是最大的，
	// ①那么就要交换，②同时使子堆保持对应相应的性质【当前结点的值大于左右儿子结点的值】
	if largest != i {
		arr[largest], arr[i] = arr[i], arr[largest]
		Heapify(arr, n, largest)
	}
}

// HeapSort 堆排序 不稳定
func HeapSort(arr []int) {
	n := len(arr)
	// 从最后一个非叶子结点开始遍历
	// 这里需要注意的是下标是否是从0开始的，不同的表示不一样，此处是从0开始的
	for i := n/2 - 1; i >= 0; i-- {
		Heapify(arr, n, i)
	}
	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0] // 与堆顶元素交换
		Heapify(arr, i, 0)
	}
}
