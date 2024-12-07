package main

import "fmt"

// 空接口类型，存储不确定类型的值
func printValue(value interface{}) {
	fmt.Println(value)
}

func main() {
	/*
		var x int = 58
		var p *int = &x
		printValue(*p)
	*/

	//var x interface{} = "Hello"
	//str, ok := x.(string)
	//if ok {
	//	println("x is a string", str)
	//} else {
	//	println("x is not a string")
	//}

	var list []interface{}
	list = append(list, "hello")
	list = append(list, 123)
	list = append(list, true)
	list = append(list, 'q')
	for i, v := range list {
		println(i, v)
	}

}
