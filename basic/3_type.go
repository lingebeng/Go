package main

import (
	. "fmt"
	"unsafe"
)

func main() {
	// 1、布尔类型
	var flag bool
	flag = false
	Println(flag)
	// 2、整型
	var x int64 = 210000000000000000 // uint
	Println(x)
	// 3、浮点型
	var y float64 = 1.20000001000000001
	Println(y)

	// 4、复数类型
	var c complex64 = 1 + 2i // 1 + 4i + 4i ^ 2 = -3 + 4i
	Println(c * c)

	// 5、数组类型，固定长度的序列
	var arr [3]int = [3]int{1, 2, 3}
	Println(arr)

	// 6、切片类型，长度不固定可自由增加
	var lst []int = []int{1, 2, 3}
	lst = append(lst, 4, 5, 6)
	Println(lst)

	// 7、Map类型
	m := map[string]int{"a": 1, "b": 2}
	Println(m)

	// 8、结构体类型
	type Person struct {
		Name string
		Age  int
	}
	p := Person{Name: "lifesaving", Age: 20}
	Println(p, p.Name, p.Age)
	// 9、字符类型
	var char byte = 'a' // byte -> uint8的别名【ascii的256个字符】   rune -> int32的别名 【表示unicode的任意字符】
	Println(char, string(char))
	var char1 rune = 'A'
	Println(unsafe.Sizeof(char), unsafe.Sizeof(char1))

	str := "Hello,World"
	Println(unsafe.Sizeof(str), str, len(str))
}
