package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

// ab相同类型，可以仅声明最后一个参数类型
func add2(a, b int) int {
	c := a + b
	return c
}

//多返回值
func exists(m map[string]string, k string) (v string, ok bool) {
	v, ok = m[k]
	return v, ok
}

func main() {
	res := add(1, 2)
	fmt.Println(res) // 3

	v, ok := exists(map[string]string{"a": "A"}, "a")
	fmt.Println(v, ok) // A True
}
