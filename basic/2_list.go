package main

import . "fmt"

func main() {
	arr := [2][2]int{
		{1, 2},
		{3, 4},
	}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			Print(arr[i][j], " ")
		}
	}

}
