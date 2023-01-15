package main

import "fmt"

// 值传递
func add2(n int) {
	n += 2
}

// 引用传递
func add2ptr(n *int) {
	*n += 2
}

func main() {
	n := 5
	add2(n)
	fmt.Println(n) // 5
	add2ptr(&n)
	fmt.Println(n) // 7
}
