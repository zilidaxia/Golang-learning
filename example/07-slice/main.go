package main

import "fmt"

// 切片是一种对数组的抽象，可以让你动态的管理数组中的元素。
// 切片是一个指向底层数组的指针，并且包含了三个属性：指针，长度和容量

func main() {
	// 使用make创建切片
	s := make([]string, 3)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("get:", s[2])   // c
	fmt.Println("len:", len(s)) // 3
	// 切片可以使用append来追加元素
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println(s) // [a b c d e f]
	// 使用copy拷贝数组
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c) // [a b c d e f]

	fmt.Println(s[2:5]) // [c d e]
	fmt.Println(s[:5])  // [a b c d e]
	fmt.Println(s[2:])  // [c d e f]

	good := []string{"g", "o", "o", "d"}
	fmt.Println(good) // [g o o d]
}
