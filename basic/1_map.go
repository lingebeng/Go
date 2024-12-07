package main

import . "fmt"

func main() {
	// 定义
	dic := map[int]int{}
	dic[1] = 0
	dic[2] = 2
	// 遍历
	for k, v := range dic {
		Println(k, v)
	}
	Println("-----------------------------------------")
	// 删除
	delete(dic, 1)
	for k, v := range dic {
		Println(k, v)
	}
	// 特殊判断
	var value int
	var ok bool
	value = dic[1]
	Println(value) // 不存在时会打印0
	value, ok = dic[1]
	if ok {
		Println(value)
	} else {
		Println("Not Exist!")
	}
}
