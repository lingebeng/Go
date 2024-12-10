package main

import (
	. "fmt"
	"sort"
	"strings"
)

func main() {
	str := "Hello,Go!"
	chars := strings.Split(str, "")
	sort.Strings(chars)
	sorted := strings.Join(chars, "")
	Println(sorted) // !,GHelloo
}
