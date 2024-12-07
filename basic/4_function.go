package main

func main() {
	var f func()
	f = func() {
		print("Hello World ")
	}
	f()
	var f1 func(a, b int) int
	f1 = func(a, b int) int {
		return a * b
	}
	println(f1(25, 2))
}
