package main

import (
	. "Go/algorithm/sort"
	"fmt"
)

func main() {
	arr := []int{1, 3, 2, 5, 4, 7, 6, 9, 8, 10}
	MergeSort(&arr, 0, len(arr)-1)
	fmt.Println(arr)
}
