package main

import (
	. "Go/algorithm/sort"
	"fmt"
)

func main() {
	a := []int{1, 4, 2, 3, 10, 9, 5, 8, 7, 6}
	QuickSort(&a, 0, len(a)-1)
	for _, x := range a {
		fmt.Printf("%d ", x)
	}
}
